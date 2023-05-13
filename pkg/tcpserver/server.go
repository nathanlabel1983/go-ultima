package tcpserver

import (
	"fmt"
	"net"
	"sync"

	"github.com/nathanlabel1983/go-ultima/pkg/packets"
)

type TCPServer struct {
	nextID  int                   // The next ID to assign to a connection
	address string                // The address to listen on
	port    int                   // The port to listen on
	conns   map[int]*net.Conn     // Map of connections, int is the connection ID
	Packets chan packets.Packeter // Channel of Packets
	mutex   sync.Mutex            // Mutex to protect the TCPServer

	//Signals
	Kill chan struct{} // Kill signal for the TCPServer
}

type TCPServerStatus struct {
	ActiveConnections int    `json:"active_connections"`
	IPAddress         string `json:"ip_address"`
	Port              int    `json:"port"`
}

func NewTCPServer(addr string, port int) *TCPServer {
	return &TCPServer{
		nextID:  0,
		address: addr,
		port:    port,
		conns:   make(map[int]*net.Conn),
		Packets: make(chan packets.Packeter, 100),
		Kill:    make(chan struct{}),
	}
}

func (s *TCPServer) Listen() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.address, s.port))
	if err != nil {
		panic(err)
	}
	defer l.Close()

	connChan := make(chan net.Conn)
	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				// Signal a nil value to signal error
				fmt.Println("TCPServer: Error accepting connection")
			} else {
				connChan <- c
			}
		}
	}()

	for {
		select {
		case <-s.Kill:
			fmt.Println("TCPServer: Kill signal received")
			return nil
		case c := <-connChan:
			// add connection
			go s.handleConnection(&c)
		}
	}
}

func (s *TCPServer) handleConnection(connection *net.Conn) {
	defer (*connection).Close()
	s.mutex.Lock()
	id := s.nextID
	s.nextID++
	s.conns[id] = connection
	s.mutex.Unlock()

	fmt.Printf("TCPServer: Connection %d accepted\n", id)

	for {
		select {
		case <-s.Kill:
			fmt.Printf("TCPServer: Connection %d closing\n", id)
			return
		default:
			// Read from connection
			packet_id := make([]byte, 1)
			_, err := (*connection).Read(packet_id)
			if err != nil {
				fmt.Printf("TCPServer: Connection %d closed\n", id)
				return
			}
			data := make([]byte, packets.PacketSizes[packet_id[0]])
			_, err = (*connection).Read(data)
			if err != nil {
				fmt.Printf("TCPServer: Connection %d closed\n", id)
				return
			}
			p := packets.NewPacket(packet_id[0], id, data)
			fmt.Printf("Packet Received: %v\n", p.GetName())
			s.Packets <- p
		}
	}
}

// GetStatus returns an int with the number of active connections
func (s *TCPServer) GetStatus() TCPServerStatus {
	status := TCPServerStatus{
		ActiveConnections: len(s.conns),
		IPAddress:         s.address,
		Port:              s.port,
	}
	return status
}

// SendPacket sends a Packeter to the connection with the specified ID
func (s *TCPServer) SendPacket(p packets.Packeter) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	conn, ok := s.conns[p.GetConnID()]
	if !ok {
		return fmt.Errorf("TCPServer: Connection %d not found", p.GetConnID())
	}
	d := append([]byte{p.GetID()}, p.GetData()...)
	_, err := (*conn).Write(d)
	if err != nil {
		return fmt.Errorf("TCPServer: Error sending packet to connection %d", p.GetConnID())
	}
	return nil
}
