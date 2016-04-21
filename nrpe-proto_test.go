package nagios

import (
//	"os"
	"testing"
	"bytes"
	"fmt"
)



func TestNRPERequest(t *testing.T) {
	var p NrpePacket
	buf := new(bytes.Buffer)
	p.SetMessage("hellfdfffffffffffffffffffffffffddddddddddddddddddddddddddddddddddddddddddo")
	p.PrepareRequest()
	err := p.Generate(buf)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())

}
