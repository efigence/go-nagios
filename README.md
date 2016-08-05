[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/efigence/go-nagios)

# go-nagios
Common data structures for nagios objects + generating them from environment

## Reading nagios status file

Structures are json-annotated so they can be dumped directly into JSON. Package uses RWMutex to lock data structure on update so it should also be used when concurrent access is needed.

```go
    file, _ := os.Open(cfg.NagiosStatusFile)
    st, err := nagios.LoadStatus(file)
    fmt.Printf("parse err: %+v\n", err)
    file.Close()
```

to update:

```go
    file, _ := os.Open(cfg.NagiosStatusFile)
    st.UpdateStatus(file)
```

Lock is only set after parsing stage so actual locked time is very short

## Generating nagios commands

Raw interface (command file is value of nagios's `command_file` config var):

```go
    cmd, err := NewCmd(`/var/nagios/nagios.cmd`)
    err = cmd.Cmd(nagios.CmdScheduleHostServiceDowntimeAll, "10", "20")
```

It is async interface and Nagios provides no sane way to check if command was actually executed so lack of error just means that write was successful, not that nagios actually did something

## Creating NRPE server/client

go-nagios have only packet serdes, to use it as client/server you need to wrap it. Simplest concurrent server looks like this:


```go
package main

import (
    "fmt"
    "net"
    "github.com/efigence/go-nagios"
)

func main() {
    sock, _ := net.Listen("tcp", ":5666")
    for {
        conn, _ := sock.Accept()
        go func(conn net.Conn) {
            req, _ := nagios.ReadNrpe(conn)
            msg, _ := req.GetMessage()
            fmt.Printf("request: %s\n", msg)
            var resp nagios.NrpePacket
            resp.SetMessage("OK")
            resp.PrepareResponse()
            // send response
            resp.Generate(conn)
            // and close
            conn.Close()

        } (conn)
    }
}

```





Currently the encryption used by NRPE by default is not suppoted in go (anonymous DH/ECDH which is vulnerable to MITM) but latest NRPE have ability to use supporte
