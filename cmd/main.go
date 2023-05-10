package main

import (
	"github.com/nathanlabel1983/go-ultima/internal/api"
	"github.com/nathanlabel1983/go-ultima/pkg/tcpserver"
)

func main() {
	s := tcpserver.NewTCPServer("127.0.0.1:2953")
	go s.Listen()
	api.StartAPI(s)
}
