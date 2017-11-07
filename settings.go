package main

import "fmt"

var (
	requestUri = ""     //请求地址
	listenPort = "8080" //监听端口
)

func SetRequestUri(uri string) {
	requestUri = uri
}

func SetListenPort(port int) {
	listenPort = fmt.Sprint(port)
}
