package nagios

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHostEncode(t *testing.T) {
	h := Host{}
	h.Hostname = "testhost"
	h.State = "UP"
	h.CheckMessage = "all green"
	Convey("TestHostEncode", t, func() {
		So(EncodeHostCheck(h), ShouldContainSubstring, "testhost;")
		So(EncodeHostCheck(h), ShouldContainSubstring, "0;")
		So(EncodeHostCheck(h), ShouldContainSubstring, "all green")
	})
}
