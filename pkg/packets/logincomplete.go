package packets

type LoginCompletePacket struct {
	packet packet
}

func (p *LoginCompletePacket) GetID() byte {
	return p.packet.id
}

func (p *LoginCompletePacket) GetName() string {
	return PacketNames[p.packet.id]
}

func (p *LoginCompletePacket) GetSize() uint16 {
	return p.packet.size
}

func (p *LoginCompletePacket) GetConnID() int {
	return p.packet.connID
}

func (p *LoginCompletePacket) GetData() []byte {
	return p.packet.data
}

func NewLoginCompletePacket(connID int) Packeter {
	return &LoginRequestPacket{
		packet: packet{
			id:     0x55,
			connID: connID,
			size:   0x0000,
			data:   []byte{},
		},
	}
}
