// Nagios notification manipulation
// see doc/notification.md for details about input format
// and http://nagios.sourceforge.net/docs/3_0/notifications.html

package nagios

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Notification types:
//
// *
type Notification struct {
	Host                 Host          `json:"hostname"`
	Service              Service       `json:"service,omitempty"`
	Type                 string        `json:"type"`
	Recipients           []string      `json:"recipients,omitempty"`
	HostState            string        `json:"host_state"`
	HostStateHard        bool          `json:"host_state_hard"`
	HostStateDuration    time.Duration `json:"host_state_duration"`
	ServiceState         string        `json:"service_state,omitempty"`
	ServiceStateHard     bool          `json:"sevice_state_hard,omitempty"`
	ServiceStateDuration time.Duration `json:"service_state_duration,omitempty"`
	IsHost               bool          `json:"is_host,omitempty"`
	IsService            bool          `json:"is_service,omitempty"`
}

func NewNotification() Notification {
	var n Notification
	n.Host = NewHost()
	n.Service = NewService()
	return n
}

// Create notification from env passed by nagios
// Intended to be used with nagios notification command
func NewNotificationFromEnv() (Notification, error) {
	n := NewNotification()
	// FIXME handle error
	var err error
	n.Host, err = NewHostFromEnv()
	n.Service, err = NewServiceFromEnv()
	n.Type = os.Getenv("NAGIOS_NOTIFICATIONTYPE")
	n.HostState = os.Getenv("NAGIOS_HOSTSTATE")
	if os.Getenv("NAGIOS_HOSTSTATETYPE") == "HARD" {
		n.HostStateHard = true
	}
	n.Recipients = strings.Split(os.Getenv("NAGIOS_NOTIFICATIONRECIPIENTS"), ",")
	h_duration, err := strconv.ParseInt(os.Getenv("NAGIOS_HOSTDURATIONSEC"), 10, 64)
	if err != nil {
		return n, err
	}
	n.HostStateDuration = time.Duration(h_duration) * time.Second

	if os.Getenv("NAGIOS_SERVICESTATE") == "" && os.Getenv("NAGIOS_HOSTSTATE") != "" {
		n.IsHost = true
	} else {
		n.IsService = true
		n.ServiceState = os.Getenv("NAGIOS_SERVICESTATE")
		s_duration, err := strconv.ParseInt(os.Getenv("NAGIOS_SERVICEDURATIONSEC"), 10, 64)
		if err != nil {
			return n, err
		}
		n.ServiceStateDuration = time.Duration(s_duration) * time.Second
		if os.Getenv("NAGIOS_SERVICESTATETYPE") == "HARD" {
			n.ServiceStateHard = true
		}
	}
	return n, err
}
