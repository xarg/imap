// Copyright 2013 The imapsrv Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the imapsrv.LICENSE file.

package main

import (
	imap "github.com/xarg/imap"
)

func main() {
	// The simplest possible server - zero config
	// It will find a free tcp port, create some temporary directories.. - just give me a server!
	//s := imap.NewServer()
	//s.Start()

	// More advanced config
	m := &imap.DummyMailstore{}

	s := imap.NewServer(
		imap.Listen("127.0.0.1:1193"),
		imap.Store(m),
	)
	s.Start()
}

// A dummy mailstore used for demonstrating the IMAP server
type DummyMailstore struct {
}

// Get mailbox information
func (m *DummyMailstore) GetMailbox(name string) (*imap.Mailbox, error) {
	return &imap.Mailbox{
		Name: "inbox",
		Id:   1,
	}, nil
}

// Get the sequence number of the first unseen message
func (m *DummyMailstore) FirstUnseen(mbox int64) (int64, error) {
	return 4, nil
}

// Get the total number of messages in an IMAP mailbox
func (m *DummyMailstore) TotalMessages(mbox int64) (int64, error) {
	return 8, nil
}

// Get the total number of unread messages in an IMAP mailbox
func (m *DummyMailstore) RecentMessages(mbox int64) (int64, error) {
	return 4, nil
}

// Get the next available uid in an IMAP mailbox
func (m *DummyMailstore) NextUid(mbox int64) (int64, error) {
	return 9, nil
}
