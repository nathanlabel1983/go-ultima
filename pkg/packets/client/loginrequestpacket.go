package client

import (
	"bytes"
)

const (
	LoginRequestPacketSize = 62
	LoginRequestPacketName = "Login Request"
	LoginRequestID         = 0x80
)

type LoginRequestPacket struct {
	connID int    // The connection ID
	data   []byte // The packet data
}

func (p *LoginRequestPacket) GetID() byte {
	return LoginRequestID
}

func (p *LoginRequestPacket) GetName() string {
	return LoginRequestPacketName
}

func (p *LoginRequestPacket) GetSize() uint16 {
	return LoginRequestPacketSize
}

func (p *LoginRequestPacket) GetConnID() int {
	return p.connID
}

func (p *LoginRequestPacket) GetData() []byte {
	return p.data
}

func NewLoginRequestPacket(connID int, data []byte) *LoginRequestPacket {
	if len(data) != LoginRequestPacketSize-1 {
		panic("LoginRequestPacket: data size is wrong")
	}
	return &LoginRequestPacket{
		connID: connID,
		data:   data,
	}
}

func (p *LoginRequestPacket) GetAccountName() string {
	// Get 30 bytes starting at 0 and convert to string, strip nulls first
	return string(bytes.Trim(p.data[0:30], "\x00"))
}

func (p *LoginRequestPacket) GetPassword() string {
	// Get 30 bytes starting at 30 and convert to string, strip nulls first
	return string(bytes.Trim(p.data[30:60], "\x00"))
}

func (p *LoginRequestPacket) GetNextLoginKey() byte {
	return p.data[60]
}
