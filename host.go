/* Host definition */
package nagios

import (
	"log"
	"os"
	"strings"
	"errors"
)

type Host struct {
	Hostname string
	Displayname string
	HostGroups []string
	Address string
	Parents *[]Host
}


func NewHost() Host {
	var h Host
	return h
}


func NewHostFromEnv() (Host, error) {
	h := NewHost()
	h.Hostname = os.Getenv("NAGIOS_HOSTNAME")
	h.Displayname = os.Getenv("NAGIOS_HOSTDISPLAYNAME")
	h.Address = os.Getenv("NAGIOS_HOSTADDRESS")
	if(os.Getenv("NAGIOS_HOSTGROUPNAMES") != "") {
		h.HostGroups = strings.Split( os.Getenv("NAGIOS_HOSTGROUPNAMES"), ",")
	}
	if (h.Hostname == "") {
		err := errors.New("Couldn't get hostname from env")
		return h, err
	}
	return h, nil
}
