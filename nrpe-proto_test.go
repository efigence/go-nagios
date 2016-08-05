package nagios

import (
	//	"os"
	"bytes"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
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
	Convey(`create packet`, t, func() {
		Convey("contains msg", func() {
			So(str, ShouldContainSubstring, testStr)
		})
		Convey("string is nul-terminated", func() {
			So(str, ShouldContainSubstring, testStr+"\000")
		})
	})
	var p2 NrpePacket
	err = p2.SetMessage(strings.Repeat("^", 65535))
	Convey("Create too big packet", t, func() {
		So(err, ShouldNotBeNil)
		So(fmt.Sprintf("%s", err), ShouldContainSubstring, "size exceed")
	})
	b,_ := p.GenerateBytes()
	nrpe, err := ReadNrpeBytes(b)
	msg, err2 := nrpe.GetMessage()
	Convey("serdes test",t,func() {
		So(err,ShouldEqual,nil)
		So(err2,ShouldEqual,nil)
		So(nrpe.Buffer,ShouldEqual,p.Buffer)
		So(msg,ShouldResemble,testStr) // resemble shows off \000 better
	})
}
