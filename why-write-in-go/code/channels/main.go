package main

import "fmt"

func main() {
	// Channels are the pipes that connect concurrent goroutines
	messages := make(chan string)

	// Send a message to our new goroutine's channel
	go func() {
		messages <- "ping"
	}()

	// Read our message from our main channel
	msg := <-messages
	fmt.Println(msg)
}
