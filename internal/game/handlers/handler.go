package handlers

import (
	"fmt"

	"github.com/nathanlabel1983/go-ultima/pkg/packets"
	"github.com/nathanlabel1983/go-ultima/pkg/packets/client"
	"github.com/nathanlabel1983/go-ultima/pkg/packets/server"
)

type PacketSender interface {
	SendPacket(p packets.Packeter) error
}

// LoginRequestPacketHandler handles the LoginRequestPacket, it will request authentication from the database.
func LoginRequestPacketHandler(s PacketSender, p packets.Packeter) {
	pkt := p.(*client.LoginRequestPacket)
	fmt.Printf("Game: %v sent %v\nUsername: %v\nPassword: %v\n", pkt.GetConnID(), pkt.GetName(), pkt.GetAccountName(), pkt.GetPassword())
	// To Implement: Handle Authentication, then if sucessful do the below.
	a := server.NewGameServerListPacket(pkt.GetConnID(), "Test")
	s.SendPacket(&a)
}
