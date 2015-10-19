package nagios

import (
	"strconv"
	"time"
	//	"os"
)

const StateOK = "OK"
const StateWarning = "WARNING"
const StateCritical = "CRITICAL"
const StateUnknown = "UNKNOWN"

var stateMapNumToName = map[string]string{
	"0": StateOK,
	"1": StateWarning,
	"2": StateCritical,
	"3": StateUnknown,
}

var stateMapNameToNum = map[string]string{
	StateOK:       "0",
	StateWarning:  "1",
	StateCritical: "2",
	StateUnknown:  "3",
}

// update fields shared by host and service
type CommonFields struct {
	Hostname            string    `json:"hostname,omitempty"`
	DisplayName         string    `json:"display_name,omitempty"`
	CheckMessage        string    `json:"check_message,omitempty"`
	State               string    `json:"service_state,omitempty"`
	PreviousState       string    `json:"previous_service_state,omitempty"`
	LastCheck           time.Time `json:"last_check,omitempty"`
	NextCheck           time.Time `json:"next_check,omitempty"`
	LastHardStateChange time.Time `json:"last_hard_state_change,omitempty"`
	LastStateChange     time.Time `json:"last_state_change,omitempty"`
	StateHard           bool      `json:"service_state_hard,omitempty"`
	Acknowledged        bool      `json:"ack,omitempty"`
	Flapping            bool      `json:"flapping,omitempty"`
	Downtime            bool      `json:"downtime,omitempty"`
}

func (c *CommonFields) UpdateCommonFromMap(m map[string]string) error {
	var err error
	c.Hostname = m["host_name"]
	c.State = stateMapNumToName[m["current_state"]]
	c.PreviousState = stateMapNumToName[m["last_hard_state"]]
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
