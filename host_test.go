package nagios

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
}

func TestHostBadEnv(t *testing.T) {
	os.Clearenv()
	_, err := NewHostFromEnv()
	if err == nil {
		t.Logf("%s", err)
		t.FailNow()
	}

}
func TestHostFromEnv(t *testing.T) {
	os.Setenv("NAGIOS_HOSTNAME", "testhost")
	os.Setenv("NAGIOS_HOSTADDRESS", "127.0.0.1")
	os.Setenv("NAGIOS_HOSTDISPLAYNAME", "long-test-host")
	os.Setenv("NAGIOS_HOSTGROUPNAMES", "hostgroup2,hostgroup1")
	host, err := NewHostFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	if host.Hostname == "" {
		t.Error("Empty host")
	}
}

func TestHostFromMap(t *testing.T) {
	m := map[string]string{
		"host_name":                     "testhost",
		"current_state":                 "1",
		"last_hard_state":               "0",
		"last_hard_state_change":        "1444749354",
		"last_state_change":             "1444749403",
		"last_check":                    "1444749433",
		"next_check":                    "1444749503",
		"scheduled_downtime_depth":      "1",
		"problem_has_been_acknowledged": "0",
		"state_type":                    "0",
		"is_flapping":                   "1",
		"plugin_output":                 "DUMMY CHECK WARNING",
	}
	tested, err := NewHostFromMap(m)
	Convey("HostFromMap", t, func() {
		So(err, ShouldEqual, nil)
		So(tested.Hostname, ShouldEqual, m["host_name"])
		So(tested.State, ShouldEqual, "WARNING")
		So(tested.StateHard, ShouldEqual, false)
		So(tested.PreviousState, ShouldEqual, "OK")
		So(tested.CheckMessage, ShouldEqual, "DUMMY CHECK WARNING")
		So(tested.LastCheck, ShouldResemble, time.Unix(1444749433, 0))
		So(tested.NextCheck, ShouldResemble, time.Unix(1444749503, 0))
		So(tested.LastStateChange, ShouldResemble, time.Unix(1444749403, 0))
		So(tested.LastHardStateChange, ShouldResemble, time.Unix(1444749354, 0))
		So(tested.Acknowledged, ShouldEqual, false)
		So(tested.Flapping, ShouldEqual, true)
	})
}

func BenchmarkHostFromEnv(b *testing.B) {
	basicEnv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewHostFromEnv()
	}
}

func BenchmarkHost(b *testing.B) {
	basicEnv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewHost()
	}
}
