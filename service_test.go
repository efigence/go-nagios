package nagios

import (
	"testing"
	"os"
)

func TestService(t *testing.T) {
	os.Setenv("NAGIOS_HOSTNAME","testhost")
	os.Setenv("NAGIOS_SERVICEDESC","test-service")
	os.Setenv("NAGIOS_SERVICEDISPLAYNAME","test-service-name")
	os.Setenv("NAGIOS_SERVICEGROUPNAMES","svcgroup1,svcgroup2")
	os.Setenv("NAGIOS_SERVICEISVOLATILE","0")
	service, err := NewServiceFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	if(service.Hostname == "") {
		t.Error("Empty host")
	}
}


func TestFailIfBadEnv(t *testing.T) {
	os.Clearenv()
	service, err := NewServiceFromEnv()
	if err == nil {
		t.Log("Should fail if env is empty")
		t.FailNow()
	}
	_ = service
}
