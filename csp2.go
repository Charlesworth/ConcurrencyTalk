package main

import "fmt"
import "time"

var Channel = make(chan string)

func main() {

	go ping()
	go pong()

	time.Sleep(10 * time.Second)
}

func ping() {
	for {
		Channel <- "Ping"

		time.Sleep(time.Second)

		msg := <-Channel
		fmt.Println(msg)
	}
}

func pong() {
	for {
		msg := <-Channel
		fmt.Println(msg)

		Channel <- "Pong"

		time.Sleep(time.Second)
	}
}
