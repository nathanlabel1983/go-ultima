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
	conns   map[int]*net.Conn     // Map of connections, int is the connection ID
	packets chan packets.Packeter // Channel of Packets
	mutex   sync.Mutex            // Mutex to protect the TCPServer

	//Signals
	Kill chan struct{} // Kill signal for the TCPServer
}

type TCPServerStatus struct {
	ActiveConnections int    `json:"active_connections"`
	IPAddress         string `json:"ip_address"`
}

func NewTCPServer(addr string) *TCPServer {
	return &TCPServer{
		nextID:  0,
		address: addr,
		conns:   make(map[int]*net.Conn),
		packets: make(chan packets.Packeter, 100),
		Kill:    make(chan struct{}),
	}
}

func (s *TCPServer) Listen() error {
	l, err := net.Listen("tcp", s.address)
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
			p := packets.NewPacket(packet_id[0], id, uint16(len(data)), data)
			s.packets <- p
		}
	}
}

// GetStatus returns an int with the number of active connections
func (s *TCPServer) GetStatus() TCPServerStatus {
	status := TCPServerStatus{
		ActiveConnections: len(s.conns),
		IPAddress:         s.address,
	}
	return status
}
