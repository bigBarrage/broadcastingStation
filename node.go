package main

import "golang.org/x/net/websocket"

type Node struct {
	IsAlive bool            //节点是否存活
	Conn    *websocket.Conn //链接
}

func (this *Node) InStation() {
	Station = append(Station, this)
}

func (this *Node) SendMessage(msg string) {
	messageChannel <- msg
}
