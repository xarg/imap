package main

import "github.com/xarg/imap"

func main() {
	// More advanced config
	m := &imap.DummyMailstore{}

	s := imap.NewServer(
		imap.Listen("127.0.0.1:1193"),
		imap.Store(m),
	)
	s.Start()
}
