package cmd

import (
	"bufio"
	"fmt"
	"log"

	"github.com/hwipl/smc-go/pkg/socket"
)

// RunClient runs the program as a client
func RunClient(host string, port int) string {
	// connect via smc
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := socket.Dial(address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Printf("Connected to server\n")

	// sent text, read reply an
	text := *text + "\n"
	fmt.Fprintf(conn, text)
	fmt.Printf("Sent %d bytes to server: %s", len(text), text)
	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes from server: %s", len(reply), reply)
	return reply
}
