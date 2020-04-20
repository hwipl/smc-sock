package cmd

import "flag"

var (
	// command line arguments
	server  = flag.Bool("s", false, "run server")
	client  = flag.Bool("c", false, "run client")
	address = flag.String("a", "127.0.0.1", "server/client address")
	port    = flag.Int("p", 50000, "server/client port")
)

// Run is the main entry point of the program
func Run() {
	// parse command line arguments
	flag.Parse()

	// run server
	if *server {
		RunServer(*address, *port)
		return
	}

	// run client
	if *client {
		RunClient(*address, *port)
		return
	}

	flag.Usage()
}
