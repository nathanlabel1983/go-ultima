package shared

import "github.com/nathanlabel1983/go-ultima/pkg/packets"

// Account Flags
const (
	AuthenticatedFlag = 1 << iota // Set if the account is authenticated
	SeededFlag                    // Set if the account is seeded
)

// Service Names
const (
	AuthServiceName = "Authentication"
	TCPServiceName  = "TCPServer"
)

// Account represents a user account
type Account struct {
	ID        string `json:"ID"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	LastLogon string `json:"LastLogon"`

	Flags  int // The account flags
	ConnID int // The connection ID
}

// GameData represents the data for a game server
type GameData struct {
	ShardName string                 // The name of the Server
	IPAddress string                 // The IP Address of the Server
	Services  map[string]ServicePort // Services that the game provides
	Accounts  map[int]Account        // Accounts that are currently connected, map by connection ID
}

// ServicePort is the interface that all services must implement
type ServicePort interface {
	Start() error
	Stop() error
}

type AuthenticationServicePort interface {
	ServicePort
	AuthAccount(username, password string) bool
}

type TCPServerServicePort interface {
	ServicePort
	GetPackets() *chan packets.Packeter
	SendPacket(p packets.Packeter) error
}

type ServicerPort interface {
	RegisterService(name string, service ServicePort)
	GetService(name string) ServicePort
	DeregisterService(name string)
}
