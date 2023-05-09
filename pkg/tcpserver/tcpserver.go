package tcpserver

import (
	"fmt"
	"net"
	"sync"

	"github.com/nathanlabel1983/go-ultima/pkg/packets"
)

type TCPServer struct {
	nextID int               // The next ID to assign to a connection
	conns  map[int]*net.Conn // Map of connections, int is the connection ID
	mutex  sync.Mutex        // Mutex to protect the TCPServer

	//Signals
	Kill chan struct{} // Kill signal for the TCPServer
}

func NewTCPServer() *TCPServer {
	return &TCPServer{
		nextID: 0,
		conns:  make(map[int]*net.Conn),
		Kill:   make(chan struct{}),
	}
}

func (s *TCPServer) Listen(address string) error {
	l, err := net.Listen("tcp", address)
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
			data := make([]byte, packets.PacketSizes[packet_id[0]])
			fmt.Println(data)
			if err != nil {
				fmt.Printf("TCPServer: Connection %d closed\n", id)
				return
			}

		}
	}

}
