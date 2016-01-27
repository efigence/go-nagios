package nagios

import (
	"strconv"
	"time"
	//	"os"
)

const StateOk = "OK"
const StateWarning = "WARNING"
const StateCritical = "CRITICAL"
const StateUnknown = "UNKNOWN"

const StateUp = "UP"
const StateDown = "DOWN"
const StateUnreachable = "UNREACHABLE"

var serviceStateMapNumToName = map[string]string{
	"0": StateOk,
	"1": StateWarning,
	"2": StateCritical,
	"3": StateUnknown,
}

var hostStateMapNumToName = map[string]string{
	"0": StateUp,
	"1": StateDown,
	"2": StateUnreachable,
}

var hostStateMapNameToNum = map[string]string{
	StateUp:          "0",
	StateDown:        "1",
	StateUnreachable: "2",
}
var serviceStateMapNameToNum = map[string]string{
	StateOk:       "0",
	StateWarning:  "1",
	StateCritical: "2",
	StateUnknown:  "3",
}

const isService = 1
const isHost = 2

// update fields shared by host and service
type CommonFields struct {
	Hostname             string    `json:"hostname,omitempty"`
	DisplayName          string    `json:"display_name,omitempty"`
	CheckMessage         string    `json:"check_message,omitempty"`
	State                string    `json:"state,omitempty"`
	PreviousState        string    `json:"previous_state,omitempty"`
	LastCheck            time.Time `json:"last_check,omitempty"`
	NextCheck            time.Time `json:"next_check,omitempty"`
	LastHardStateChange  time.Time `json:"last_hard_state_change,omitempty"`
	LastStateChange      time.Time `json:"last_state_change,omitempty"`
	StateHard            bool      `json:"state_hard"`
	Acknowledged         bool      `json:"ack"`
	Flapping             bool      `json:"flapping"`
	Downtime             bool      `json:"downtime"`
	NotificationsEnabled bool      `json:"notifications_enabled"`
}

// update common fields of service/host
// dataType should be either isService or isHost
func (c *CommonFields) UpdateCommonFromMap(m map[string]string, dataType int) error {
	var err error
	c.Hostname = m["host_name"]
	if dataType == isHost {
		c.State = hostStateMapNumToName[m["current_state"]]
		c.PreviousState = hostStateMapNumToName[m["last_hard_state"]]
	} else if dataType == isService {
		c.State = serviceStateMapNumToName[m["current_state"]]
		c.PreviousState = serviceStateMapNumToName[m["last_hard_state"]]
	} else {
		panic("unkown type passed, pass either isHost or isService const")
	}
	c.CheckMessage = m["plugin_output"]

	i, err := strconv.ParseInt(m["last_hard_state_change"], 10, 64)
	if err != nil {
		return err
	}
	c.LastHardStateChange = time.Unix(i, 0)

	i, err = strconv.ParseInt(m["last_state_change"], 10, 64)
	if err != nil {
		return err
	}
	c.LastStateChange = time.Unix(i, 0)

	i, err = strconv.ParseInt(m["last_check"], 10, 64)
	if err != nil {
		return err
	}
	c.LastCheck = time.Unix(i, 0)

	i, err = strconv.ParseInt(m["next_check"], 10, 64)
	if err != nil {
		return err
	}
	c.NextCheck = time.Unix(i, 0)

	i, err = strconv.ParseInt(m["state_type"], 10, 64)
	if err != nil {
		return err
	}
	if i > 0 {
		c.StateHard = true
	} else {
		c.StateHard = false
	}

	i, err = strconv.ParseInt(m["notifications_enabled"], 10, 64)
	if err != nil {
		return err
	}
	if i > 0 {
		c.NotificationsEnabled = true
	} else {
		c.NotificationsEnabled = false
	}

	i, err = strconv.ParseInt(m["scheduled_downtime_depth"], 10, 64)
	if err != nil {
		return err
	}
	// yes it can go negative
	// nagios developers do not know why but they duct-taped it eventually
	// http://sourceforge.net/p/nagios/nagioscore/ci/cdb47278535b7957726328cc30a18eb5753521f5/
	if i > 0 {
		c.Downtime = true
	} else {
		c.Downtime = false
	}

	i, err = strconv.ParseInt(m["problem_has_been_acknowledged"], 10, 64)
	if err != nil {
		return err
	}
	if i > 0 {
		c.Acknowledged = true
	} else {
		c.Acknowledged = false
	}
	i, err = strconv.ParseInt(m["is_flapping"], 10, 64)
	if err != nil {
		return err
	}
	if i > 0 {
		c.Flapping = true
	} else {
		c.Flapping = false
	}

	return err
}
