package nagios

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNagios(t *testing.T) {
	service, err := NewServiceFromArgs([]string{"hostname","servicename","1"})
	Convey("Service from args", t, func() {
		So( err, ShouldNotBeNil)
	})

	service, err = NewServiceFromArgs([]string{"hostname","servicename","1","test"})
	Convey("Service from args", t, func() {
		So( err, ShouldBeNil)
	})
	service.UpdateStatus("OK","test")
	Convey("", t, func() {
		So( err, ShouldBeNil)
	})
	service.UpdateStatus("OK","test")
	Convey("", t, func() {
		So( err, ShouldBeNil)
	})
	number_fields := []string{
		"last_hard_state_change",
		"last_state_change",
		"last_check",
		"next_check",
		"state_type",
		"notifications_enabled",
		"scheduled_downtime_depth",
		"problem_has_been_acknowledged",
		"is_flapping",
	}

	for _,fieldname := range number_fields {
		err = service.UpdateCommonFromMap(map[string]string{fieldname: "not a number"}, isService)
		Convey(fieldname + " bad number", t, func() {
			So(err, ShouldNotBeNil)
		})
		err = service.UpdateCommonFromMap(map[string]string{fieldname: "1"}, isService)
		Convey(fieldname + " good number", t, func() {
			So(err, ShouldBeNil)
		})
	}
}
