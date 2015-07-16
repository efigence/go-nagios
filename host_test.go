package nagios

import (
	"testing"
	"os"
)

func TestMain(t *testing.T) {
}

func TestHostBadEnv(t *testing.T) {
	os.Clearenv()
	_, err := NewHostFromEnv()
	if err == nil {
		t.Logf("%s", err)
		t.FailNow()
	}

}
func TestHostFromEnv(t *testing.T) {
	os.Setenv("NAGIOS_HOSTNAME","testhost")
	os.Setenv("NAGIOS_HOSTADDRESS","127.0.0.1")
	os.Setenv("NAGIOS_HOSTDISPLAYNAME","long-test-host")
	os.Setenv("NAGIOS_HOSTGROUPNAMES","hostgroup2,hostgroup1")
	host, err := NewHostFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	if(host.Hostname == "") {
		t.Error("Empty host")
	}
}

func BenchmarkHostFromEnv(b *testing.B) {
	basicEnv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewHostFromEnv()
	}
}

func BenchmarkHost(b *testing.B) {
	basicEnv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_= NewHost()
	}
}
