package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("something ith i " + fmt.Sprint(i))
	defer fmt.Println("asdasdasdasdasdasd" + fmt.Sprint(i))
	fmt.Println("something happens")
	i++
	return
}
