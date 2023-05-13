package client

import "encoding/binary"

const (
	SelectServerPacketSize = 3
	SelectServerPacketName = "Select Server"
	SelectServerPacketID   = 0xA0
)

type SelectServerPacket struct {
	connID int    // The connection ID
	data   []byte // The packet data
}

func (p *SelectServerPacket) GetID() byte {
	return SelectServerPacketID
}

func (p *SelectServerPacket) GetName() string {
	return SelectServerPacketName
}

func (p *SelectServerPacket) GetSize() uint16 {
	return SelectServerPacketSize
}

func (p *SelectServerPacket) GetConnID() int {
	return p.connID
}

func (p *SelectServerPacket) GetData() []byte {
	return p.data
}

func NewSelectServerPacket(connID int, data []byte) *SelectServerPacket {
	if len(data) != SelectServerPacketSize-1 {
		panic("SelectServerPacket: data size is wrong")
	}
	return &SelectServerPacket{
		connID: connID,
		data:   data,
	}
}

func (p *SelectServerPacket) GetShard() uint16 {
	// return uint16 (Which is 2 bytes)
	return binary.BigEndian.Uint16(p.data[0:2])
}
