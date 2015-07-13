package nagios

import (
	"testing"
)

func TestService(t *testing.T) {
	service, err := NewServiceFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	if(service.Hostname == "") {
		t.Error("Empty host")
	}
}
