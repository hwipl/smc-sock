package cmd

import (
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestRunClient(t *testing.T) {
	var want, got string

	// start a server
	l := prepareServer("127.0.0.1", 0)
	go runServerLoop(l)

	// get address and port for client
	addr := strings.Split(l.Addr().String(), ":")
	port, err := strconv.Atoi(addr[1])
	if err != nil {
		log.Fatal(err)
	}

	// run test with default text message
	want = *text + "\n"
	got = RunClient(addr[0], port)
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}

	// run test with custom text message
	*text = "Hello again!!!"
	want = *text + "\n"
	got = RunClient(addr[0], port)
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}

	// run test with empty text message
	*text = ""
	want = *text + "\n"
	got = RunClient(addr[0], port)
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}
}
