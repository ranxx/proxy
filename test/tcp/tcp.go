package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":12444")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("新连接", conn.RemoteAddr().String())
		_ = conn
	}
}
