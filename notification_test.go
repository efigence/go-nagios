package nagios

import (
	"os"
	"testing"
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

func TestNotificationSoft(t *testing.T) {
	basicEnv()
	os.Setenv("NAGIOS_HOSTSTATETYPE", "SOFT")
	os.Setenv("NAGIOS_SERVICESTATETYPE", "SOFT")
	n, err := NewNotificationFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	if n.HostStateHard {
		t.Error("Host state should be soft")
	}
	if n.ServiceStateHard {
		t.Error("Service state should be soft")
	}
}

func TestNotificationHost(t *testing.T) {
	basicEnv()
	os.Unsetenv("NAGIOS_SERVICESTATE")
	n := testNotification(t)

	if !n.IsHost {
		t.Error("Host not detected as host")
	}
	if n.IsService {
		t.Errorf("Host detected as service")
	}
}

func TestBadTime(t *testing.T) {
	basicEnv()
	os.Setenv("NAGIOS_SERVICEDURATIONSEC", "kitten")
	n, err := NewNotificationFromEnv()
	if err == nil {
		t.Logf("Should detect bad time")
		t.FailNow()
	}
	_ = n
}

func testNotification(t *testing.T) Notification {
	n, err := NewNotificationFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	return n
}

func BenchmarkNotificationFromEnv(b *testing.B) {
	basicEnv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewNotificationFromEnv()
	}
}

func BenchmarkNotification(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewNotification()
	}
}
