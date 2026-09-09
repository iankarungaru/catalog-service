package main

import (
	"fmt"
	"time"

)
func sayThing(thing string) {
	fmt.Println(thing)
}

func main() {
	go sayThing("Hello from a goroutine")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main function finishing")

	fmt.Println("---")

	for i := 1; i <= 3; i++ {
		go sayThing(fmt.Sprintf("Task %d", i))
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println("All done")
}