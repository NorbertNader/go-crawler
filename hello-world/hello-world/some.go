package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

const addr = "localhost:5555"

var pattern *regexp.Regexp

func main() {
	pattern = regexp.MustCompile("/|:|\\.|<|>|\\*|\\|")
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(addr, nil)
	fmt.Println("error" + err.Error())
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request from", r.RemoteAddr, r.Method)

	if r.Method == http.MethodPost {
		scanner := bufio.NewScanner(r.Body)
		total := 0
		jobs := make(chan string, 100)
		results := make(chan string, 100)

		for i := 1; i <= 3; i++ {
			go fetchWorker(i, jobs, results)
		}

		for scanner.Scan() {
			jobs <- scanner.Text()
			total++
		}
		close(jobs)
		fmt.Println("total jobs " + fmt.Sprintf("%d", total))

		if total > 0 {
			for j := 1; j <= total; j++ {
				<-results
			}
		}

		if err := scanner.Err(); err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	} else if r.Method == http.MethodGet {
		fmt.Println("hello to", r.RemoteAddr)
		w.Header().Set("Content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprintf(w, "Hello World\n")

		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func safePath(url string) string {
	return pattern.ReplaceAllString(url, "_")
}

func fetchWorker (id int, jobs <-chan string, results chan<-string) {
	fmt.Println("worker", id, "started")
	for j := range jobs {
		fmt.Println("Worer", id, "starting for job", j)
		result := fetch(j)
		fmt.Println("Worker", id, "finished for job", j)
		fmt.Println(result)
		results <- result
	}
}

func fetch (url string) (result string) {
	response, err := http.Get(url)

	fmt.Println("I am in fetch")

	if err != nil {
		result = "Failed to fetch due to error " + url + " " + err.Error()
	} else if response.StatusCode != http.StatusOK {
		result = "Failed to fetch due to status code " + url + " " + fmt.Sprintf("%d", response.StatusCode)
	} else {
		defer response.Body.Close()
		fmt.Println("the path is: " + safePath(url))
		fname := "tmp/" + safePath(url)
		f, err := os.Create(fname)

		if err != nil {
			result = "Error creating file " + fname + " " + err.Error()
		} else {
			defer f.Close()
			_, err := io.Copy(f, response.Body)

			if err != nil {
				result = "Error copying response of " + url + " " + err.Error()
			} else {
				result = "Fetched " + url + " as " + fname
			}
		}
	}
	return
}
