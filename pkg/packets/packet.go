package packets

import (
	"fmt"

	"github.com/nathanlabel1983/go-ultima/pkg/packets/client"
)

type Packeter interface {
	GetID() byte
	GetName() string
	GetSize() uint16
	GetConnID() int
	GetData() []byte
}

type Packet struct {
	ID     byte   // The 1 byte packet ID
	ConnID int    // The connection ID
	Size   uint16 // The 2 byte packet size
	Data   []byte // The packet data
}

// NewPacket returns a Packeter based on the packet ID
func NewPacket(id byte, connID int, data []byte) Packeter {
	switch id {
	case client.LoginSeedPacketID:
		return client.NewLoginSeedPacket(connID, data)
	case client.LoginRequestID:
		return client.NewLoginRequestPacket(connID, data)
	case client.SelectServerPacketID:
		return client.NewSelectServerPacket(connID, data)
	default:
		fmt.Printf("Client sent packet: %x\n", id)
		panic("Packet Not implemented.\n")
	}
}
