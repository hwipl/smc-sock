package cmd

import "flag"

// variable definitions
var (
	server  bool   = false
	client  bool   = false
	address string = "127.0.0.1"
	port    int    = 50000
)

// Run is the main entry point of the program
func Run() {
	// parse command line arguments
	flag.BoolVar(&server, "s", server, "run server")
	flag.BoolVar(&client, "c", client, "run client")
	flag.StringVar(&address, "a", address, "server/client address")
	flag.IntVar(&port, "p", port, "server/client port")

	flag.Parse()

	// run server
	if server {
		RunServer(address, port)
		return
	}

	// run client
	if client {
		RunClient(address, port)
		return
	}

	flag.Usage()
}
