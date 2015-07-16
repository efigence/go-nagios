package nagios

import (
	"testing"
	"os"
)


func TestService(t *testing.T) {
	basicEnv()
	service, err := NewServiceFromEnv()
	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
	if(service.Hostname == "") {
		t.Error("Empty host")
	}
}

func TestServiceNoDesc(t *testing.T) {
	basicEnv()
	os.Unsetenv("NAGIOS_SERVICEDESC")
	_, err := NewServiceFromEnv()
	if err == nil {
		t.Logf("%s", err)
		t.FailNow()
	}

}


func TestFailIfBadEnv(t *testing.T) {
	os.Clearenv()
	os.Setenv("NAGIOS_SERVICEISVOLATILE","1")
	service, err := NewServiceFromEnv()
	if err == nil {
		t.Log("Should fail if env is empty")
		t.FailNow()
	}
	_ = service
}


func BenchmarkServiceFromEnv(b *testing.B) {
	basicEnv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewServiceFromEnv()
	}
}
func BenchmarkService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewService()
	}
}
