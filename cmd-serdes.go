package nagios

import (
	"fmt"
	"strings"
)

// Decode host check string
func DecodeHostCheck(s string) (Host, error) {
	var parts []string
	var h Host
	if strings.HasPrefix(s, CmdProcessHostCheckResult) {
		parts = strings.SplitN(s, ";", 4)
		if len(parts) < 4 {
			return h, fmt.Errorf("Decode error, not enough parts after splitting [%s]", s)
		}
		parts = parts[1:]
	} else {
		parts = strings.SplitN(s, ";", 3)
		if len(parts) < 3 {
			return h, fmt.Errorf("Decode error, not enough parts after splitting [%s]", s)
		}
	}
	h.Hostname = parts[0]
	if val, ok := hostStateMapNumToName[parts[1]]; ok {
		h.State = val
	} else {
		h.State = StateUnknown
	}
	h.CheckMessage = parts[2]
	return h, nil
}

// Encode host status into host check string (without PASSIVE_HOST_CHECK_RESULT header)
func EncodeHostCheck(h Host) string {
	fmt.Printf("%+v", hostStateMapNameToNum)
	return strings.Join([]string{
		h.Hostname,
		hostStateMapNameToNum[h.State],
		h.CheckMessage,
	},";")

}
