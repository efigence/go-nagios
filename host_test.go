package nagios

import (
	"testing"
)

func TestHost(t *testing.T) {
	host, err := NewHostFromEnv()
	if err != nil {
		t.Log(": %s", err)
		t.FailNow()
	}
	if(host.Hostname == "") {
		t.Log("Empty host")
		t.FailNow()
	}
}
