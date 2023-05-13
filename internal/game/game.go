package game

import (
	"github.com/nathanlabel1983/go-ultima/data"
	"github.com/nathanlabel1983/go-ultima/internal/api"
	"github.com/nathanlabel1983/go-ultima/internal/game/handlers"
	"github.com/nathanlabel1983/go-ultima/internal/services"
	"github.com/nathanlabel1983/go-ultima/pkg/packets/client"
	"github.com/nathanlabel1983/go-ultima/pkg/tcpserver"
)

const (
	AuthServiceName = "Authentication"
)

type Game struct {
	tcpserver *tcpserver.TCPServer
	shardName string
	config    data.Config
	services  map[string]Service // Services that the game provides
	// Signals
	Kill chan struct{}
}

func NewGame() *Game {
	g := Game{
		services: make(map[string]Service),
		Kill:     make(chan struct{}),
	}
	g.config = data.LoadConfiguration("\\data\\configuration.json") // Load all config data
	// Register all services

	// register AuthService
	fp, _ := g.config.GetConfigPath("Accounts")
	g.RegisterService(AuthServiceName, services.NewAuthService(fp))

	// now start all services
	for _, s := range g.services {
		go s.Start()
	}
	// Now configure the TCP Server for the game
	g.tcpserver = tcpserver.NewTCPServer(g.config.Server.IPAddress, g.config.Server.Port)
	g.shardName = g.config.Game.ShardName
	return &g
}

func (g *Game) RegisterService(name string, service Service) {
	g.services[name] = service
}

func (g *Game) GetService(name string) Service {
	return g.services[name]
}

func (g *Game) DeregisterService(name string) {
	delete(g.services, name)
}

func (g *Game) HandlePackets() {
	for {
		select {
		case <-g.Kill:
			return
		case p := <-g.tcpserver.Packets:
			// Use type switch to handle packets
			switch p.(type) {
			case *client.LoginRequestPacket:
				handlers.LoginRequestPacketHandler(g.tcpserver, p)
			}
		}
	}
}

func (g *Game) Start() {
	go g.tcpserver.Listen()
	go g.HandlePackets()
	api.StartAPI(g.tcpserver)
}

func (g *Game) GetConfig() data.Config {
	return g.config
}
