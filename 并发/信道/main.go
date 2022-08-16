package main

import "fmt"

func main() {
	var channel chan string = make(chan string)
	var msg string

	go message("hello go", channel)
	// 取消息
	msg = <-channel
	fmt.Println(msg)
}

func message(msg string, msgChannel chan string) {
	// 存消息
	msgChannel <- msg
}
