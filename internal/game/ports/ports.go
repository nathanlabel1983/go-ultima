package ports

import (
	"github.com/nathanlabel1983/go-ultima/pkg/packets"
	"github.com/nathanlabel1983/go-ultima/pkg/shared"
)

type ServicePort interface {
	Start() error
	Stop() error
}

type AuthenticationServicePort interface {
	ServicePort
	AuthAccount(username, password string) (shared.Account, error)
}

type TCPServerServicePort interface {
	ServicePort
	GetPackets() *chan packets.Packeter
}

type PacketSenderPort interface {
	SendPacket(p packets.Packeter) error
}

type ServicerPort interface {
	RegisterService(name string, service ServicePort)
	GetService(name string) ServicePort
	DeregisterService(name string)
}
