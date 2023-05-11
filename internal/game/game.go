package game

import (
	"fmt"

	"github.com/nathanlabel1983/go-ultima/internal/api"
	"github.com/nathanlabel1983/go-ultima/pkg/packets"
	"github.com/nathanlabel1983/go-ultima/pkg/tcpserver"
)

type Game struct {
	tcpserver *tcpserver.TCPServer

	// Signals
	Kill chan struct{}
}

func NewGame(tcpserver *tcpserver.TCPServer) *Game {
	return &Game{
		tcpserver: tcpserver,
		Kill:      make(chan struct{}),
	}
}

func (g *Game) HandlePackets() {
	for {
		select {
		case <-g.Kill:
			return
		case p := <-g.tcpserver.Packets:
			// Use type switch to handle packets
			switch p.(type) {
			case *packets.LoginRequestPacket:
				// Handle LoginRequestPacket
				// Cast p to *packets.LoginRequestPacket
				loginRequestPacket := p.(*packets.LoginRequestPacket)
				fmt.Printf("Game: Connection  {%v} sent LoginRequestPacket. Username: %v, Password: %v", loginRequestPacket.GetConnID(), loginRequestPacket.GetAccountName(), loginRequestPacket.GetPassword())
				// Testing: Send back accepted packet
				a := packets.NewLoginCompletePacket(loginRequestPacket.GetConnID())
				g.tcpserver.SendPacket(a)
			}
		}
	}
}

func (g *Game) Start() {
	go g.tcpserver.Listen()
	go g.HandlePackets()
	api.StartAPI(g.tcpserver)
}
