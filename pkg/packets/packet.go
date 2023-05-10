package packets

type Packeter interface {
	GetID() byte
	GetName() string
	GetSize() uint16
}

type packet struct {
	id     byte   // The 1 byte packet ID
	connID int    // The connection ID
	size   uint16 // The 2 byte packet size
	data   []byte // The packet data
}

// NewPacket returns a Packeter based on the packet ID
func NewPacket(id byte, connID int, size uint16, data []byte) Packeter {
	p := packet{id: id, connID: connID, size: size, data: data}
	switch id {
	case 0xEF:
		return newLoginSeedPacket(p)
	default:
		return nil
	}
}
