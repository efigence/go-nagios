package nagios

import (
	//	"os"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"bytes"
	"fmt"
)



func TestNRPERequest(t *testing.T) {
	var p NrpePacket
	buf := new(bytes.Buffer)
	testStr := "=hellfdfffffddddddddddddz"
	p.SetMessage(testStr)
	p.PrepareRequest()
	err := p.Generate(buf)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	str := fmt.Sprintf("%s", buf.Bytes())
	Convey (`create packet`,t, func() {
		Convey("contains msg",func() {
			So(str,ShouldContainSubstring,testStr)
		})
		Convey("string is nul-terminated",func() {
			So(str,ShouldContainSubstring,testStr +"\000")
		})
	})

}
