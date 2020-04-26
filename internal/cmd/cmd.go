package cmd

import "flag"

var (
	// command line arguments
	server  = flag.Bool("s", false, "run as server")
	client  = flag.Bool("c", false, "run as client")
	address = flag.String("a", "127.0.0.1:50000",
		"listen on (server) or connect to (client) `address`")
	text = flag.String("t", "Hello, world!", "transfer `text`")
)

// Run is the main entry point of the program
func Run() {
	// parse command line arguments
	flag.Parse()

	// run server
	if *server {
		RunServer(*address)
		return
	}

	// run client
	if *client {
		RunClient(*address)
		return
	}

	flag.Usage()
}
