package main

var messageChannel chan string

func init() {
	messageChannel = make(chan string, 1024)
}
