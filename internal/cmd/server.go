package cmd

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/hwipl/smc-go/pkg/socket"
)

// prepareServer prepares the listener for the server
func prepareServer(address string, port int) net.Listener {
	// listen for smc connections
	l, err := socket.Listen(address, port)
	if err != nil {
		log.Fatal(err)
	}
	return l
}

// runServerLoop runs the endless server loop
func runServerLoop(l net.Listener) {
	defer l.Close()
	// accept new connections from listener and handle them
	for {
		// accept new connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// handle new connection: send data back to client
		go func(c net.Conn) {
			fmt.Printf("New client connection\n")
			written, err := io.Copy(c, c)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Echoed %d bytes to client\n", written)
			c.Close()
		}(conn)
	}
}

// RunServer runs the program as a server
func RunServer(address string, port int) {
	l := prepareServer(address, port)
	runServerLoop(l)
}
