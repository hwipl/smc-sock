package cmd

import (
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestRunClient(t *testing.T) {
	// start a server
	l := prepareServer("127.0.0.1", 0)
	go runServerLoop(l)

	// run client
	addr := strings.Split(l.Addr().String(), ":")
	port, err := strconv.Atoi(addr[1])
	if err != nil {
		log.Fatal(err)
	}
	RunClient(addr[0], port)
}
