package packets

import (
	"bytes"
)

type LoginRequestPacket struct {
	packet packet
}

func (p *LoginRequestPacket) GetID() byte {
	return p.packet.id
}

func (p *LoginRequestPacket) GetName() string {
	return PacketNames[p.packet.id]
}

func (p *LoginRequestPacket) GetSize() uint16 {
	return p.packet.size
}

func (p *LoginRequestPacket) GetConnID() int {
	return p.packet.connID
}

func (p *LoginRequestPacket) GetData() []byte {
	return p.packet.data
}

func newLoginRequestPacket(p packet) Packeter {
	return &LoginRequestPacket{packet: p}
}

func (p *LoginRequestPacket) GetAccountName() string {
	// Get 30 bytes starting at 0 and convert to string, strip nulls first
	return string(bytes.Trim(p.packet.data[0:30], "\x00"))
}

func (p *LoginRequestPacket) GetPassword() string {
	// Get 30 bytes starting at 30 and convert to string, strip nulls first
	return string(bytes.Trim(p.packet.data[30:60], "\x00"))
}

func (p *LoginRequestPacket) GetNextLoginKey() byte {
	return p.packet.data[60]
}
