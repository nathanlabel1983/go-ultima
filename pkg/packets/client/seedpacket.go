package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	LoginSeedPacketSize = 21
	LoginSeedPacketName = "KR/2D Client Login/Seed"
	LoginSeedPacketID   = 0xEF
)

type LoginSeedPacket struct {
	connID int    // The connection ID
	data   []byte // The packet data
}

func (p *LoginSeedPacket) GetID() byte {
	return LoginSeedPacketID
}

func (p *LoginSeedPacket) GetName() string {
	return LoginSeedPacketName
}

func (p *LoginSeedPacket) GetSize() uint16 {
	return LoginSeedPacketSize
}

func (p *LoginSeedPacket) GetConnID() int {
	return p.connID
}

func (p *LoginSeedPacket) GetData() []byte {
	return p.data
}

func NewLoginSeedPacket(connID int, data []byte) *LoginSeedPacket {
	if len(data) != LoginSeedPacketSize-1 {
		panic("KR/2D Client Login/Seed: data size is wrong")
	}
	return &LoginSeedPacket{
		connID: connID,
		data:   data,
	}
}

// GetSeed returns the seed, usually also the client IP
func (p *LoginSeedPacket) GetSeed() string {
	return fmt.Sprintf("%d.%d.%d.%d", p.data[4], p.data[5], p.data[6], p.data[7])
}

func (p *LoginSeedPacket) GetMajor() (int, error) {
	return getIntVal(p.data[4:8])
}

func (p *LoginSeedPacket) GetMinor() (int, error) {
	return getIntVal(p.data[8:12])
}

func (p *LoginSeedPacket) GetRevision() (int, error) {
	return getIntVal(p.data[12:16])
}

func (p *LoginSeedPacket) GetPrototype() (int, error) {
	return getIntVal(p.data[16:20])
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
