/* status.dat handling */

package nagios

import (
	"bufio"
	"io"
	"regexp"
	"sync"
)

var blockStartRegex = regexp.MustCompile(`^(contactstatus|hoststatus|info|programstatus|servicecomment|servicestatus)\s\{`)
var blockEndRegex = regexp.MustCompile(`^\s}$`)
var kvRegex = regexp.MustCompile(`^\s*(\S+?)=(.*)$`)

type Status struct {
	Host    map[string]Host               `json:"host,omitempty"`
	Service map[string]map[string]Service `json:"service,omitempty"`
	Summary Summary                       `json:"summary,omitempty"`
	sync.RWMutex
}

func LoadStatus(r io.Reader) (Status, error) {
	var status Status
	var err error
	scanner := bufio.NewScanner(r)
	status.Host = make(map[string]Host)
	status.Service = make(map[string]map[string]Service)
	block_type := ""
	block_content := make(map[string]string)
	for scanner.Scan() {
		t := scanner.Text()
		match := blockStartRegex.FindStringSubmatch(t)
		if len(match) > 1 {
			block_type = match[1]
			continue
		}
		match = kvRegex.FindStringSubmatch(t)
		if len(match) > 1 {
			block_content[match[1]] = match[2]
			continue
		}
		if blockEndRegex.MatchString(t) {
			if block_type == "servicestatus" {
				s, err := NewServiceFromMap(block_content)
				if err == nil {
					if status.Service[s.Hostname] == nil {
						status.Service[s.Hostname] = make(map[string]Service)
					}
					status.Service[s.Hostname][s.Description] = s

				} else {
					return status, err
				}

			} else if block_type == "hoststatus" {
				h, err := NewHostFromMap(block_content)
				if err == nil {
					status.Host[h.Hostname] = h
				} else {
					return status, err
				}

			}
			//end of block summary ,cleanup vars
			block_type = ""
			block_content = make(map[string]string)
			continue
		}
	}
	status.Summary.UpdateHost(status.Host)
	status.Summary.UpdateService(status.Service)
	return status, err
}

func (s *Status) UpdateStatus(r io.Reader) error {
	status, err := LoadStatus(r)
	s.Lock()
	s.Host = status.Host
	s.Service = status.Service
	s.Unlock()
	return err
}
