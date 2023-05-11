package main

import (
	"github.com/nathanlabel1983/go-ultima/internal/game"
	"github.com/nathanlabel1983/go-ultima/pkg/tcpserver"
)

func main() {
	s := tcpserver.NewTCPServer("127.0.0.1:2953")
	g := game.NewGame(s)
	g.Start()
}
