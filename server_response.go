// Copyright 2014 The imapsrv Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the imapsrv.LICENSE file.
package imap

import (
	"bufio"
)

// An IMAP serverResponse
type serverResponse struct {
	// The tag of the command that this is the serverResponse for
	tag string
	// The machine readable condition
	condition string
	// A human readable message
	message string
	// Untagged output lines
	untagged []string
	// Should the connection be closed after the serverResponse has been sent?
	closeConnection bool
}

// Create a serverResponse
func createResponse(tag string, condition string, message string) *serverResponse {
	return &serverResponse{
		tag:       tag,
		condition: condition,
		message:   message,
		untagged:  make([]string, 0, 4),
	}
}

// Create a OK serverResponse
func ok(tag string, message string) *serverResponse {
	return createResponse(tag, "OK", message)
}

// Create an BAD serverResponse
func bad(tag string, message string) *serverResponse {
	return createResponse(tag, "BAD", message)
}

// Create a NO serverResponse
func no(tag string, message string) *serverResponse {
	return createResponse(tag, "NO", message)
}

// Write an untagged fatal serverResponse
func fatalResponse(w *bufio.Writer, err error) {
	resp := createResponse("*", "BYE", err.Error())
	resp.closeConnection = true
	resp.write(w)
}

// Add an untagged line to a serverResponse
func (r *serverResponse) extra(line string) *serverResponse {
	r.untagged = append(r.untagged, line)
	return r
}

// Mark that a serverResponse should close the connection
func (r *serverResponse) shouldClose() *serverResponse {
	r.closeConnection = true
	return r
}

// Write a serverResponse to the given writer
func (r *serverResponse) write(w *bufio.Writer) error {

	// Write untagged lines
	for _, line := range r.untagged {
		_, err := w.WriteString("* " + line + "\r\n")
		if err != nil {
			return err
		}
	}

	_, err := w.WriteString(r.tag + " " + r.condition + " " + r.message + "\r\n")
	if err != nil {
		return err
	}

	// Flush the serverResponse
	w.Flush()
	return nil
}
