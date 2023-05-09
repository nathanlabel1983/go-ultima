package packets

type LoginSeedPacket struct {
	packet packet
}

func (p *LoginSeedPacket) GetID() byte {
	return p.packet.id
}

func (p *LoginSeedPacket) GetName() string {
	return PacketNames[p.packet.id]
}

func (p *LoginSeedPacket) GetSize() uint16 {
	return p.packet.size
}
