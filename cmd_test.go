package nagios

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCmd(t *testing.T) {
	filename := "t-data/nagios.cmd"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	f.Close()

	cmd, err := NewCmd(filename)
	Convey("new command", t, func() {
		So(err, ShouldEqual, nil)
	})

	err = cmd.Cmd(CmdScheduleHostServiceDowntimeAll, "a", "b")
	Convey("Run command via Cmd", t, func() {
		So(err, ShouldEqual, nil)
	})
	err = cmd.Send(CmdScheduleServiceDowntime, "c", "d", "e")
	Convey("Run command via Cmd", t, func() {
		So(err, ShouldEqual, nil)
	})
	cmd.Close()
	data, err := ioutil.ReadFile(filename)
	Convey("Test output", t, func() {
		So(err, ShouldEqual, nil)
		So(string(data), ShouldContainSubstring, "SCHEDULE_HOST_SVC_DOWNTIME;a;b\n")
		So(string(data), ShouldContainSubstring, "SCHEDULE_SVC_DOWNTIME;c;d;e\n")
	})

}
