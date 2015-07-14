package nagios

import (
	"testing"
	"os"
)

func TestNotificationBadEnv(t *testing.T) {
	os.Clearenv()
	n, err := NewNotificationFromEnv()
	if err == nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	_ = n
}

func TestNotification(t *testing.T) {
	os.Setenv("NAGIOS_HOSTNAME","testhost")
	os.Setenv("NAGIOS_SERVICEDESC","test-service")
	os.Setenv("NAGIOS_HOSTADDRESS","127.0.0.1")
	os.Setenv("NAGIOS_HOSTDISPLAYNAME","long-test-host")
	n, err := NewNotificationFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	_ = n
}
