package server

import "encoding/binary"

const (
	GameServerListPacketSize = 46
	GameServerListPacketName = "Game Server List"
	GameServerListPacketID   = 0xA8
)

type GameServerListPacket struct {
	connID int    // The connection ID
	data   []byte // The packet data
}

func (p *GameServerListPacket) GetID() byte {
	return GameServerListPacketID
}

func (p *GameServerListPacket) GetName() string {
	return GameServerListPacketName
}

func (p *GameServerListPacket) GetSize() uint16 {
	return GameServerListPacketSize
}

func (p *GameServerListPacket) GetConnID() int {
	return p.connID
}

func (p *GameServerListPacket) GetData() []byte {
	return p.data
}

func NewGameServerListPacket(connID int, serverName string) *GameServerListPacket {
	p := GameServerListPacket{
		connID: connID,
		data:   make([]byte, GameServerListPacketSize-1),
	}
	// Set first 2 bytes to the size of the packet (uint16)
	binary.BigEndian.PutUint16(p.data[0:2], p.GetSize())
	// Set the 3rd byte to 0xCC
	p.data[2] = 0xCC
	// Now the number of servers which is 1 as uint16
	binary.BigEndian.PutUint16(p.data[3:5], 1)
	// Server index as uint16
	binary.BigEndian.PutUint16(p.data[5:7], 0)
	// Server name as 32 bytes
	copy(p.data[7:39], serverName)
	// now a byte representing how full server is (set to 0)
	p.data[39] = 0
	// Now a byte representing timezone
	p.data[40] = 0
	// Now 4 bytes representing IP address to ping (reversed)
	p.data[41] = 1
	p.data[42] = 0
	p.data[43] = 0
	p.data[44] = 127
	return &p
}
