package nagios

import (
	//	"os"
	"fmt"
	"net"
)

func ExampleNrpePacket() {
	sock, _ := net.Listen("tcp", ":5666")
	for {
		conn, _ := sock.Accept()
		go func(conn net.Conn) {
			req, _ := ReadNrpe(conn)
			msg, _ := req.GetMessage()
			fmt.Printf("request: %s\n", msg)
			var resp NrpePacket
			resp.SetMessage("OK")
			resp.PrepareResponse()
			// send response
			resp.Generate(conn)
			// and close
			conn.Close()

		}(conn)
	}
}
