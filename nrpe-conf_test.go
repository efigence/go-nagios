package nagios

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestNrpeConf(t *testing.T) {
	f, ferr := os.Open("t-data/nrpe.cfg")
	cfg, err := ParseNrpeConfig(f)
	Convey("Load file", t, func() {
		So(ferr, ShouldEqual, nil)
		So(err, ShouldEqual, nil)
		So(cfg.Config["debug"], ShouldEqual, "0")
		So(cfg.Command["check_cron"], ShouldEqual, "/usr/lib64/nagios/plugins/check_procs -a crond -c 1:10 -t 180")
	})
}
