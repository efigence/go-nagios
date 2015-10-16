/* Service definition */
package nagios

import (
	"errors"
	"os"
	//	"strconv"
	"strings"
	//	"time"
)

type Service struct {
	CommonFields
	Description   string   `json:"description,omitempty"`
	ServiceGroups []string `json:"servicegroups,omitempty"` // list of service groups this service belongs to
	Volatile      bool     `json:"volatile,omitempty"`
	Contacts      []string `json:"contacts,omitempty"`      // list of service contacts
	ContactGroups []string `json:"contactgroups,omitempty"` // list of service contact groups
}

func NewService() Service {
	var s Service
	return s
}

// create service from nagios env variables
func NewServiceFromEnv() (Service, error) {
	s := NewService()
	s.Hostname = os.Getenv("NAGIOS_HOSTNAME")
	s.Description = os.Getenv("NAGIOS_SERVICEDESC")
	s.DisplayName = os.Getenv("NAGIOS_SERVICEDISPLAYNAME")
	if os.Getenv("NAGIOS_SERVICEISVOLATILE") == "1" {
		s.Volatile = true
	} else {
		s.Volatile = false // nagios default for service
	}
	if os.Getenv("NAGIOS_SERVICEGROUPNAMES") != "" {
		s.ServiceGroups = strings.Split(os.Getenv("NAGIOS_SERVICEGROUPNAMES"), ",")
	}
	if s.Hostname == "" {
		err := errors.New("Couldn't get service hostname from env")
		return s, err
	}
	if s.Description == "" {
		err := errors.New("Couldn't get service description from env")
		return s, err
	}
	return s, nil
}

// Generate service data from key/value pairs in "status.dat" format
func NewServiceFromMap(m map[string]string) (Service, error) {
	//	s := NewService()
	var s Service
	var err error
	err = s.UpdateCommonFromMap(m)
	if err != nil {
		return s, err
	}
	s.Description = m["service_description"]
	return s, err
}
