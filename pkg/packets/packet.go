package packets

type Packeter interface {
	GetID() byte
	GetName() string
	GetSize() uint16
}

type packet struct {
	id   byte
	size uint16
	data []byte
}
