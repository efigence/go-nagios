/* Host definition */
package nagios

import (
	"errors"
	"os"
	"strings"
)

type Host struct {
	Hostname    string   `json:"hostname"`
	Displayname string   `json:"display_name"`
	HostGroups  []string `json:"hostgroups"`
	Address     string   `json:"address"`
	Parents     *[]Host  `json:"parents"`
}

func NewHost() Host {
	var h Host
	return h
}


// Create host object from nagios env variables
func NewHostFromEnv() (Host, error) {
	h := NewHost()
	h.Hostname = os.Getenv("NAGIOS_HOSTNAME")
	h.Displayname = os.Getenv("NAGIOS_HOSTDISPLAYNAME")
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
