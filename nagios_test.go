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
	bad := make(map[string]string,len(number_fields))
	good := make(map[string]string,len(number_fields))
	for _,fieldname := range number_fields {
		good[fieldname] = "1"
		bad[fieldname] = "1"
	}
	err = service.UpdateCommonFromMap(good, isService)
	Convey("all good number s", t, func() {
		So(err, ShouldBeNil)
	})
	for _,fieldname := range number_fields {
		bad[fieldname] = "not a number"
		err = service.UpdateCommonFromMap(map[string]string{fieldname: "1"}, isService)
		Convey(fieldname + " bad number", t, func() {
			So(err, ShouldNotBeNil)
		})
		bad[fieldname] = "1"
	}
}
