package nagios

import (
	"testing"
)

func TestHost(t *testing.T) {
	host, err := NewHostFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	if(host.Hostname == "") {
		t.Error("Empty host")
	}
}