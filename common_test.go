package nagios

import (
	"os"
)

func basicEnv () {
	os.Setenv("NAGIOS_HOSTNAME","testhost")
	os.Setenv("NAGIOS_SERVICEDESC","test-service")
	os.Setenv("NAGIOS_SERVICEDISPLAYNAME","test-service-name")
	os.Setenv("NAGIOS_SERVICEGROUPNAMES","svcgroup1,svcgroup2")
	os.Setenv("NAGIOS_SERVICEISVOLATILE","0")
}
