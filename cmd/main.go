package main

import "github.com/nathanlabel1983/go-ultima/pkg/tcpserver"

func main() {
	s := tcpserver.NewTCPServer()
	s.Listen("127.0.0.1")
}
