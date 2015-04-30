package main

import "fmt"

var Channel = make(chan string)

func main() {

	go helloWorld()

	msg := <-Channel
	fmt.Println(msg)
}

func helloWorld() {
	Channel <- "Hello from a go routine"
}
