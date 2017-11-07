package main

var messageChannel chan []byte

func init() {
	messageChannel = make(chan []byte, 1024)
	go sendMessage()
}
