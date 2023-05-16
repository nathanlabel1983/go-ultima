package game

import (
	"github.com/nathanlabel1983/go-ultima/data"

	"github.com/nathanlabel1983/go-ultima/pkg/packets/client"
	"github.com/nathanlabel1983/go-ultima/pkg/services/authentication"
	"github.com/nathanlabel1983/go-ultima/pkg/services/tcpserver"
	"github.com/nathanlabel1983/go-ultima/pkg/shared"
)

const (
	AuthServiceName = "Authentication"
	TCPServiceName  = "TCPServer"
)

type Game struct {
	shared.GameData
	config data.Config
	// Signals
	Kill chan struct{}
}

func NewGame() *Game {
	g := Game{
		GameData: shared.GameData{
			Accounts: make(map[int]shared.Account),
			Services: make(map[string]shared.ServicePort),
		},
	}

	data.LoadConfiguration("\\data\\configuration.json") // Load all config data
	// Register all services

	// Register TCP Server
	g.RegisterService(TCPServiceName,
		tcpserver.NewTCPServer(
			g.config.Server.IPAddress,
			g.config.Server.Port,
		),
	)

	// register AuthService
	fp, _ := g.config.GetConfigPath("Accounts")
	g.RegisterService(AuthServiceName, authentication.NewAuthService(fp))

	g.ShardName = g.config.Game.ShardName
	g.IPAddress = g.config.Server.IPAddress
	return &g
}

func (g *Game) RegisterService(name string, service shared.ServicePort) {
	g.Services[name] = service
}

func (g *Game) GetService(name string) shared.ServicePort {
	return g.Services[name]
}

func (g *Game) DeregisterService(name string) {
	delete(g.Services, name)
}

func (g *Game) Start() error {
	for _, s := range g.Services {
		go s.Start()
	}
	go g.HandlePackets()
	return nil
}

func (g *Game) Stop() error {
	for _, s := range g.Services {
		s.Stop()
	}
	g.Kill <- struct{}{}
	return nil
}

func (g *Game) GetConfig() data.Config {
	return g.config
}

func (g *Game) HandlePackets() {
	pkts := g.Services[TCPServiceName].(*tcpserver.TCPServer).GetPackets()
	for {
		select {
		case <-g.Kill:
			return
		case p := <-*pkts:
			// Use type switch to handle packets
			switch p.(type) {
			case *client.LoginSeedPacket:
				LoginSeedPacketHandler(g.GameData, p)
			}
		}
	}
}
