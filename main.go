package main

import (
	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/internal/bootstrap"
)

func init() {
	config.ParseFile("config.yaml")
}

func main() {

	// go func() {
	// 	// 启动ws
	// 	net.Listen("tcp", ":1")
	// }

	bootstrap.Start()
}
