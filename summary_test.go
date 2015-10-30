package nagios

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestSummary(t *testing.T) {
	file, err := os.Open("t-data/status.dat")
	if err != nil {
			t.Logf("%s", err)
		t.FailNow()
	}
	s, err := LoadStatus(file)
	_ = err
	Convey("get summary from status", t, func() {
		So(err, ShouldEqual, nil)
		Convey("host stats",func() {
			So(s.Summary.HostCount.All, ShouldEqual, 4)
			So(s.Summary.HostCount.Up, ShouldEqual, 2)
			So(s.Summary.HostCount.Down, ShouldEqual, 1)
			So(s.Summary.HostCount.Unreachable, ShouldEqual, 1)
		})
		Convey("service stats",func() {
			So(s.Summary.ServiceCount.All, ShouldEqual, 7)
			So(s.Summary.ServiceCount.Ok, ShouldEqual, 7)
			So(s.Summary.ServiceCount.Warning, ShouldEqual, 7)
			So(s.Summary.ServiceCount.Critical, ShouldEqual, 7)
		})
	})

}
