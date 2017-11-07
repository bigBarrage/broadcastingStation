package main

import "github.com/gorilla/websocket"

func sendMessage() {
	for {
		msg := <-messageChannel

		for _, node := range Station {
			w, err := node.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				continue
			}
			w.Write(msg)
			w.Close()
		}
	}
}
