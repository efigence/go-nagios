package nagios

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
	"fmt"
)

func TestServiceFromEnv(t *testing.T) {
	basicEnv()
	service, err := NewServiceFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	if service.Hostname == "" {
		t.Error("Empty host")
	}
	Convey("ServiceFromEnv", t, func() {
		So(service.Hostname, ShouldEqual, "testhost")
		So(service.Description, ShouldEqual, "test-service")
		So(service.State, ShouldEqual, "CRITICAL")
		So(service.StateHard, ShouldEqual, true)
	})

}

func TestServiceMarshal(t *testing.T) {
	basicEnv()
	service, err := NewServiceFromEnv()
	fmt.Printf("%+v", service)
	Convey("ServiceMarshal", t, func() {
		So(err,ShouldBeNil)
		So(service.MarshalCmd(), ShouldEqual, "testhost;test-service;2;DUMMY CHECK CRITICAL")
	})
}

func TestServiceNoDesc(t *testing.T) {
	basicEnv()
	os.Unsetenv("NAGIOS_SERVICEDESC")
	_, err := NewServiceFromEnv()
	if err == nil {
		t.Logf("%s", err)
		t.FailNow()
	}

}

func TestFailIfBadEnv(t *testing.T) {
	os.Clearenv()
	os.Setenv("NAGIOS_SERVICEISVOLATILE", "1")
	service, err := NewServiceFromEnv()
	if err == nil {
		t.Log("Should fail if env is empty")
		t.FailNow()
	}
	_ = service
}

func TestServiceFromMap(t *testing.T) {
	m := map[string]string{
		"host_name":                     "testhost",
		"service_description":           "test service",
		"current_state":                 "1",
		"last_hard_state":               "0",
		"last_hard_state_change":        "1444749354",
		"last_state_change":             "1444749403",
		"last_check":                    "1444749433",
		"next_check":                    "1444749503",
		"scheduled_downtime_depth":      "0",
		"state_type":                    "1",
		"problem_has_been_acknowledged": "1",
		"is_flapping":                   "0",
		"notifications_enabled":         "1",
		"plugin_output":                 "DUMMY CHECK WARNING",
	}
	tested, err := NewServiceFromMap(m)
	Convey("ServiceFromMap", t, func() {
		So(err, ShouldEqual, nil)
		So(tested.Hostname, ShouldEqual, m["host_name"])
		So(tested.Description, ShouldEqual, m["service_description"])
		So(tested.State, ShouldEqual, "WARNING")
		So(tested.StateHard, ShouldEqual, true)
		So(tested.PreviousState, ShouldEqual, "OK")
		So(tested.LastCheck, ShouldResemble, time.Unix(1444749433, 0))
		So(tested.NextCheck, ShouldResemble, time.Unix(1444749503, 0))
		So(tested.LastStateChange, ShouldResemble, time.Unix(1444749403, 0))
		So(tested.LastHardStateChange, ShouldResemble, time.Unix(1444749354, 0))
		So(tested.Acknowledged, ShouldEqual, true)
		So(tested.NotificationsEnabled, ShouldEqual, true)
		So(tested.Flapping, ShouldEqual, false)
	})

}

func BenchmarkServiceFromEnv(b *testing.B) {
	basicEnv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewServiceFromEnv()
	}
}
func BenchmarkService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewService()
	}
}
