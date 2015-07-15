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
	basicEnv()
	n, err := NewNotificationFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	_ = n
}
