package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.Create("tmp/some.txt")

	if err != nil {
		fmt.Println(err.Error() + " we hade some shit happen")
	} else {
		defer f.Close()
		fmt.Println("file created bitch!")
	}
}
