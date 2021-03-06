package nagios

import (
	"fmt"
	"strings"
)


func DecodeNagiosCmd(cmd string) (name string, args []string, err error) {
	parts := strings.Split(cmd, ";")
	// dumb but check if it is nagios cmd (all types of command are upcase and LOOKING_LIKE_THIS)
	if len(parts) >= 2 && strings.ToUpper(parts[0]) == parts[0] && strings.Contains(parts[0],"_") {
		return parts[0],parts[1:],nil
	} else if len(parts) >= 2 {
		return "",nil,fmt.Errorf("command [%s] doesn't look like NAGIOS_CMD", parts[0])
	} else {
		return "",nil,fmt.Errorf("can't parse [%s] into nagios cmd", name)
	}

}


// Decode host check string
func DecodeHostCheck(check string) (Host, error) {
	var parts []string
	if cmd, args, err := DecodeNagiosCmd(check); err == nil && cmd == CmdProcessHostCheckResult {
		if len(args) < 3 {
			return Host{}, fmt.Errorf("Decode error, not enough parts after splitting [%s]", check)
		}
		parts = args
	} else if err == nil {
		return Host{}, fmt.Errorf("Expected host check, got [%s]",cmd)
	} else {
		parts = strings.SplitN(check, ";", 3)
		if len(parts) < 3 {
			return Host{}, fmt.Errorf("Decode error, not enough parts after splitting [%s]", check)
		}
	}
	return NewHostFromArgs(parts)
}

// Encode host status into host check string (without PASSIVE_HOST_CHECK_RESULT header)
func EncodeHostCheck(h Host) string {
	return strings.Join([]string{
		h.Hostname,
		hostStateMapNameToNum[h.State],
		h.CheckMessage,
	},";")

}


func DecodeServiceCheck(check string) (Service, error) {
	var parts []string
	if cmd, args, err := DecodeNagiosCmd(check); err == nil && cmd == CmdProcessServiceCheckResult {
		if len(args) < 4 {
			return Service{}, fmt.Errorf("Decode error, not enough parts after splitting [%s]", check)
		}
		parts = args
	} else if err == nil {
		return Service{}, fmt.Errorf("Expected service check, got [%s]",cmd)
	} else {
		parts = strings.SplitN(check, ";", 4)
		if len(parts) < 4 {
			return Service{}, fmt.Errorf("Decode error, not enough parts after splitting [%s]", check)
		}
	}
	return NewServiceFromArgs(parts)
}

func EncodeServiceCheck(s Service) string {
	return strings.Join([]string{
		s.Hostname,
		s.Description,
		serviceStateMapNameToNum[s.State],
		s.CheckMessage,
	},";")

}