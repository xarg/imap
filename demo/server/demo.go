package main

import (
	imap "github.com/xarg/imap"
)

func main() {
	// The simplest possible server - zero config
	// It will find a free tcp port, create some temporary directories.. - just give me a server!
	s := imap.NewServer()
	s.Start()
}
