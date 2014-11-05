package imap
============

[![Build Status](https://travis-ci.org/xarg/imap.svg?branch=master)](https://travis-ci.org/xarg/imap)

IMAP client and server implementation in Go.

This project is the result of an unification effort of 2 projects:

* https://github.com/mxk/go-imap - IMAP client written in Go
* https://github.com/alienscience/imapsrv - IMAP server written in Go


The motivation behind this merge is code reuse, easier testing and easier contribution.

Installation
------------

```
go get github.com/xarg/imap
```


Client
------


Server
------

On it's own, the server, is not functional because it requires a mail storage, an authentication backend, etc..
It provides generic interfaces around mail storage and authentication. It must be used in conjunction with drivers.

See storage drivers.


However, it offers a few dummy interfaces as an inspiration or used for testing.


The simplest way to start a server:


```
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
```

You can try it out:

```
go run $GOPATH/src/github.com/xarg/imap/demo/server/demo.go
```


A more advanced example:

```
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

```


Docs
----

http://godoc.org/github.com/xarg/imap

License
-------

BSD
