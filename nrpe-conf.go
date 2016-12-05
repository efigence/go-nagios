package nagios

import (
	"bufio"
	"io"
	"regexp"
)

var nrpeCommandLineRe = regexp.MustCompile(`^\s*command\[(.+?)\]\s*=\s*(.+)\s*`)
var nrpeConfigLineRe = regexp.MustCompile(`^\s*(\S+)\s*=\s*(.+)$`)
var stripCommentRe = regexp.MustCompile(`\s*#.*`)

type NrpeConfig struct {
	Config  map[string]string
	Command map[string]string
}

func ParseNrpeConfig(data io.Reader) (cfg NrpeConfig, err error) {
	scan := bufio.NewScanner(data)
	cfg.Config = make(map[string]string)
	cfg.Command = make(map[string]string)
	for scan.Scan() {
		line := stripCommentRe.ReplaceAllString(scan.Text(), ``)
		matches := nrpeCommandLineRe.FindStringSubmatch(line)
		if len(matches) > 2 {
			cfg.Command[matches[1]] = matches[2]
			continue
		}
		matches = nrpeConfigLineRe.FindStringSubmatch(line)
		if len(matches) > 2 {
			cfg.Config[matches[1]] = matches[2]
			continue
		}
	}
	return cfg, err
}
