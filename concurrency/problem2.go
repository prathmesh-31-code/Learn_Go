/*
Problem 3: Send and Receive
Create a goroutine that sends "Hello Go" to a channel

Receive and print it in main
*/
package concurrency

func ChannelSend(ch1 chan string) {
	ch1 <- "Hello Go"

}

/*
package main

import (
	"fmt"
	"learn_go/concurrency"
)

func main() {
	ch1 := make(chan string)

	go concurrency.ChannelSend(ch1)

	msg := <-ch1
	fmt.Println("Message: ", msg)
}

*/
