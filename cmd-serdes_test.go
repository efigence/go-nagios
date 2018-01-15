package nagios

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHostEncode(t *testing.T) {
	h := Host{}
	h.Hostname = "testhost"
	h.State = StateUp
	h.CheckMessage = "all green"
	Convey("TestHostEncode", t, func() {
		So(EncodeHostCheck(h), ShouldContainSubstring, "testhost;")
		So(EncodeHostCheck(h), ShouldContainSubstring, "0;")
		So(EncodeHostCheck(h), ShouldContainSubstring, "all green")
	})
}

func TestHostDecode(t *testing.T) {
	h := CmdProcessHostCheckResult + ";testhost2;1;bad"
	host,err := DecodeHostCheck(h)
	Convey("TestHostDecode1 - host down", t, func() {
		So(err, ShouldBeNil)
		So(host.Hostname, ShouldEqual, "testhost2")
		So(host.CheckMessage, ShouldEqual, "bad")
		So(host.State, ShouldEqual, StateDown)
	})

	h = "testhost2;1"
	host,err = DecodeHostCheck(h)
	Convey("TestHostDecode - bad data", t, func() {
		So(err, ShouldNotBeNil)
	})

	h = "PROCESS_HOST_CHECK_RESULT;testhost3;2;unk"
	host,err = DecodeHostCheck(h)
	Convey("TestHostDecode - host unreachable", t, func() {
		So(err, ShouldBeNil)
		So(host.Hostname, ShouldEqual, "testhost3")
		So(host.CheckMessage, ShouldEqual, "unk")
		So(host.State, ShouldEqual, StateUnreachable)
	})
}

func TestServiceEncode(t *testing.T) {
	s := Service{}
	s.Hostname = "testhost1"
	s.Description = "testservice1"
	s.CheckMessage = "check ok"
	s.State = StateOk

	Convey("TestHostEncode", t, func() {
		So(EncodeServiceCheck(s), ShouldContainSubstring, "testhost1;")
		So(EncodeServiceCheck(s), ShouldContainSubstring, "testservice1;")
		So(EncodeServiceCheck(s), ShouldContainSubstring, "0;")
		So(EncodeServiceCheck(s), ShouldContainSubstring, "check ok")
	})

}

func TestServiceDecode(t *testing.T) {
	service, err := DecodeServiceCheck(CmdProcessHostCheckResult + ";testhost2;testservice2;2;bad")
	Convey("TestServiceDecode - service critical", t, func() {
		So(err, ShouldBeNil)
		So(service.Hostname, ShouldEqual, "testhost2")
		So(service.Description, ShouldEqual, "testservice2")
		So(service.CheckMessage, ShouldEqual, "bad")
		So(service.State, ShouldEqual, StateCritical)
	})
	service,err = DecodeServiceCheck("testhost2;testserv;1")
	Convey("TestServiceDecode - bad data", t, func() {
		So(err, ShouldNotBeNil)
	})
	service, err = DecodeServiceCheck("testhost3;testservice3;0;kk")
	Convey("TestServiceDecode - service ok", t, func() {
		So(err, ShouldBeNil)
		So(service.Hostname, ShouldEqual, "testhost3")
		So(service.Description, ShouldEqual, "testservice3")
		So(service.CheckMessage, ShouldEqual, "kk")
		So(service.State, ShouldEqual, StateOk)
	})
}