/* Host definition */
package nagios

import (
	"errors"
	"os"
	"strings"
	"fmt"
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
// NewHostFromArgs creates host from nagios cmd args slice
func NewHostFromArgs(args []string) (Host, error) {
	var h Host
	h.Hostname = args[0]
	if val, ok := hostStateMapNumToName[args[1]]; ok {
		h.State = val
	} else {
		return h, fmt.Errorf("invalid host state [%s]", args[1])
	}
	h.CheckMessage = args[2]
	return h,nil
}
// MarshalCmd marshals host data in nagios cmd-compatible format (';'-separated fields)
// it is mainly designed to be used with Command.Send, like this:
//
//       cmd.Send(nagios.CmdProcessServiceCheckResult,service.MarshalCmd())

func (h *Host)MarshalCmd() string{
	h.RLock()
	defer h.RUnlock()
	return EncodeHostCheck(*h)
}
