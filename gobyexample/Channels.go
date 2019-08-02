package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "second"

	}()

	msg := <-messages
	msg = <-messages
	fmt.Println(msg)
}
