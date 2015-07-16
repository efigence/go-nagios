package nagios

import (
	"os"
	"testing"
)

func basicEnv () {
	os.Setenv("NAGIOS_HOSTNAME","testhost")
	os.Setenv("NAGIOS_HOSTDURATIONSEC","300")
	os.Setenv("NAGIOS_HOSTSTATETYPE","HARD")
	os.Setenv("NAGIOS_HOSTSTATE","UP")
	os.Setenv("NAGIOS_SERVICESTATE","CRITICAL")
	os.Setenv("NAGIOS_SERVICESTATETYPE","HARD")
	os.Setenv("NAGIOS_SERVICEDESC","test-service")
	os.Setenv("NAGIOS_SERVICEDISPLAYNAME","test-service-name")
	os.Setenv("NAGIOS_SERVICEGROUPNAMES","svcgroup1,svcgroup2")
	os.Setenv("NAGIOS_SERVICEISVOLATILE","0")
	os.Setenv("NAGIOS_SERVICEDURATION","0d 0h 5m 0s")
    os.Setenv("NAGIOS_SERVICEDURATIONSEC","300")

}


func BenchmarkSetenv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.Setenv("NAGIOS_SERVICEDURATIONSEC","300")
	}
}

func BenchmarkGetenv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = os.Getenv("NAGIOS_SERVICEDURATIONSEC")
	}
}
