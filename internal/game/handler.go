package game

import (
	"fmt"

	"github.com/nathanlabel1983/go-ultima/pkg/packets"
	"github.com/nathanlabel1983/go-ultima/pkg/packets/client"
	"github.com/nathanlabel1983/go-ultima/pkg/packets/server"
	"github.com/nathanlabel1983/go-ultima/pkg/shared"
)

type PacketSender interface {
	SendPacket(p packets.Packeter) error
}

// LoginRequestPacketHandler handles the LoginRequestPacket, it will request authentication from the database.
func LoginRequestPacketHandler(d shared.GameData, p packets.Packeter) {
	pkt := p.(*client.LoginRequestPacket)
	fmt.Printf("Game: %v sent %v\nUsername: %v\nPassword: %v\n", pkt.GetConnID(), pkt.GetName(), pkt.GetAccountName(), pkt.GetPassword())
	// First check if the account exists, then check that it has been seeded
	// If it has been seeded, then check the password, if it is correct, then
	// set the account to authenticated. Finally Send a message to the client

	a, ok := d.Accounts[pkt.GetConnID()]
	if !ok {
		fmt.Printf("Account does not exist for connection: %v\n", pkt.GetConnID())
		return
	}
	if a.Flags&shared.SeededFlag == 0 {
		fmt.Printf("Account has not been seeded for connection: %v\n", pkt.GetConnID())
		return
	}
	if a.Flags&shared.AuthenticatedFlag != 0 {
		fmt.Printf("Account is already authenticated for connection: %v\n", pkt.GetConnID())
		return
	}

	// Get the Authentication Service
	var as shared.AuthenticationServicePort = d.Services[AuthServiceName].(shared.AuthenticationServicePort)
	accID, err := as.AuthAccount(pkt.GetAccountName(), pkt.GetPassword())
	if err != nil {
		fmt.Printf("Error authenticating account: %v\n", err)
		return
	}
	a.Flags |= shared.AuthenticatedFlag // Set the authenticated flag
	a.Username = pkt.GetAccountName()
	a.Password = pkt.GetPassword()
	a.ID = accID
	d.Accounts[pkt.GetConnID()] = a
	// Now that its authenticated, LoginRequestPacketHandler should always send the GameServer list
	var s shared.TCPServerServicePort = d.Services[TCPServiceName].(shared.TCPServerServicePort)
	var serverPkt packets.Packeter = server.NewGameServerListPacket(a.ConnID, "Test")
	go s.SendPacket(serverPkt)

}

func LoginSeedPacketHandler(d shared.GameData, p packets.Packeter) (*shared.Account, error) {
	pkt := p.(*client.LoginSeedPacket)
	fmt.Printf("Game: %v sent %v\n", pkt.GetConnID(), pkt.GetName())
	_, ok := d.Accounts[pkt.GetConnID()]
	if ok {
		fmt.Printf("Account already exists for connection: %v\n", pkt.GetConnID())
		return nil, fmt.Errorf("Account already exists for connection: %v", pkt.GetConnID())
	} else {
		a := shared.Account{
			ConnID: pkt.GetConnID(),
			Flags:  shared.SeededFlag,
		}
		d.Accounts[pkt.GetConnID()] = a
		return &a, nil
	}
}
