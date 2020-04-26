package cmd

import (
	"testing"
)

func TestRunClient(t *testing.T) {
	var want, got string

	// start a server
	l := prepareServer("127.0.0.1:0")
	go runServerLoop(l)

	// run test with default text message
	want = *text + "\n"
	got = RunClient(l.Addr().String())
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}

	// run test with custom text message
	*text = "Hello again!!!"
	want = *text + "\n"
	got = RunClient(l.Addr().String())
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}

	// run test with empty text message
	*text = ""
	want = *text + "\n"
	got = RunClient(l.Addr().String())
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}
}
