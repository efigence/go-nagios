/* Service definition */
package nagios

import (
	"log"
	"os"
	"strings"
	"errors"
)


type Service struct {
	Hostname string
	Description string
	DisplayName string
	ServiceGroups []string // list of service groups this service belongs to
	Volatile bool
	Contacts []string // list of service contacts
	ContactGroups []string // list of service contact groups
}


func NewService() Service {
	var s Service
	return s
}

func NewServiceFromEnv() (Service, error) {
	s := NewService()
	s.Hostname = os.Getenv("NAGIOS_HOSTNAME")
	s.Description = os.Getenv("NAGIOS_SERVICEDESC")
	s.DisplayName = os.Getenv("NAGIOS_SERVICEDISPLAYNAME")
	if (os.Getenv("NAGIOS_SERVICEISVOLATILE") == "1") {
		s.Volatile = true
	} else {
		s.Volatile = false // nagios default for service
	}
	if(os.Getenv("NAGIOS_SERVICEGROUPNAMES") != "") {
		s.ServiceGroups = strings.Split( os.Getenv("NAGIOS_SERVICEGROUPNAMES"), ",")
	}
	if (s.Hostname == "") {
		err := errors.New("Couldn't get service hostname from env")
		return s, err
	}
	if (s.Description == "") {
		err := errors.New("Couldn't get service description from env")
		return s, err
	}
	return s, nil
}
