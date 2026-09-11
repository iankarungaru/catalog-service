package main

import "fmt"

func main() {
	ch := make(chan string)
	message := <-ch
	fmt.Println(message)

	
}