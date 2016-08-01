package nagios

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestCmd(t *testing.T) {
	f, err := os.Create("t-data/nagios.cmd")
	if err != nil {
		panic(err)
	}
	f.Close()

	cmd, err := NewCmd(`t-data/nagios.cmd`)
	_ = cmd
	Convey("N", t, func() {
		So(err, ShouldEqual, nil)
	})
	err = cmd.Cmd(CmdScheduleHostServiceDowntimeAll, "10", "20")
}
