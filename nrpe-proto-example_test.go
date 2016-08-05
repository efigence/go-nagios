package nagios

import (
	//	"os"
	"fmt"
	"net"
	"github.com/efigence/go-nagios"
)


func ExampleNrpePacket() {
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
