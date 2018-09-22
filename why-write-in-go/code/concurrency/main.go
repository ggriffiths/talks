package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Direct function call
	f("direct")

	// Called in a Goroutine (logical thread)
	go f("goroutine")
}
