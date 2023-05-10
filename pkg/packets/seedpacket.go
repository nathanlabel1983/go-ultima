package packets

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

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

func newLoginSeedPacket(p packet) Packeter {
	return &LoginSeedPacket{packet: p}
}

// GetSeed returns the seed, usually also the client IP
func (p *LoginSeedPacket) GetSeed() string {
	return fmt.Sprintf("%d.%d.%d.%d", p.packet.data[4], p.packet.data[5], p.packet.data[6], p.packet.data[7])
}

func (p *LoginSeedPacket) GetMajor() (int, error) {
	return getIntVal(p.packet.data[4:8])
}

func (p *LoginSeedPacket) GetMinor() (int, error) {
	return getIntVal(p.packet.data[8:12])
}

func (p *LoginSeedPacket) GetRevision() (int, error) {
	return getIntVal(p.packet.data[12:16])
}

func (p *LoginSeedPacket) GetPrototype() (int, error) {
	return getIntVal(p.packet.data[16:20])
}

// GetVersion returns the version of the client
func (p *LoginSeedPacket) GetVersion() (string, error) {
	major, err := p.GetMajor()
	if err != nil {
		return "", err
	}
	minor, err := p.GetMinor()
	if err != nil {
		return "", err
	}
	rev, err := p.GetRevision()
	if err != nil {
		return "", err
	}
	proto, err := p.GetPrototype()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d", major, minor, rev, proto), nil
}

// GetIntVal converts a byte array to an int
func getIntVal(val []byte) (int, error) {
	b := bytes.NewReader(val)
	var i int
	err := binary.Read(b, binary.BigEndian, &i)
	if err != nil {
		return 0, err
	}
	return i, nil
}
