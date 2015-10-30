/* Host definition */
package nagios

import (
	"errors"
	"os"
	"strings"
)

type Host struct {
	CommonFields
	HostGroups []string `json:"hostgroups"`
	Address    string   `json:"address"`
	Parents    *[]Host  `json:"parents"`
}

func NewHost() Host {
	var h Host
	return h
}

// Create host object from nagios env variables
func NewHostFromEnv() (Host, error) {
	h := NewHost()
	h.Hostname = os.Getenv("NAGIOS_HOSTNAME")
	h.DisplayName = os.Getenv("NAGIOS_HOSTDISPLAYNAME")
	h.Address = os.Getenv("NAGIOS_HOSTADDRESS")
	if os.Getenv("NAGIOS_HOSTGROUPNAMES") != "" {
		h.HostGroups = strings.Split(os.Getenv("NAGIOS_HOSTGROUPNAMES"), ",")
	}
	if h.Hostname == "" {
		err := errors.New("Couldn't get hostname from env")
		return h, err
	}
	return h, nil
}

func NewHostFromMap(m map[string]string) (Host, error) {
	h := NewHost()
	var err error
	h.UpdateCommonFromMap(m, isHost)
	return h, err
}
