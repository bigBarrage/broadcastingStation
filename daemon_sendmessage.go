package main

func sendMessage() {
	for {
		msg := <-messageChannel

		for _, node := range Station {
			node.Conn.Write(msg)
		}
	}
}
