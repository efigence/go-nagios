/* Service definition */
package nagios

import (
	"errors"
	"os"
	//	"strconv"
	"strings"
	//	"time"
	"fmt"
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
	s.State = os.Getenv("NAGIOS_SERVICESTATE")
	if os.Getenv("NAGIOS_SERVICESTATETYPE") == "HARD" {
		s.StateHard = true
	}
	s.CheckMessage = os.Getenv("NAGIOS_SERVICEOUTPUT")
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
	err = s.UpdateCommonFromMap(m, isService)
	if err != nil {
		return s, err
	}
	s.Description = m["service_description"]
	return s, err
}

func NewServiceFromArgs(args []string) (Service, error) {
	var s Service
	if len(args) < 4 {
		return s, fmt.Errorf("too little arguments, should have host, service, state, check message")
	}
	s.Hostname = args[0]
	s.Description = args[1]
	if val, ok := serviceStateMapNumToName[args[2]]; ok {
		s.State = val
	} else {
		s.State = StateUnknown
	}
	s.CheckMessage = args[3]
	return s, nil
}

// MarshalCmd marshals service data in nagios cmd-compatible format (';'-separated fields)
// it is mainly designed to be used with Command.Send, like this:
//
//       cmd.Send(nagios.CmdProcessServiceCheckResult,service.MarshalCmd())

func (s *Service)MarshalCmd() string{
	s.RLock()
	defer s.RUnlock()
	return EncodeServiceCheck(*s)
}