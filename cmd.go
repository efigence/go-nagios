package nagios

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// https://old.nagios.org/developerinfo/externalcommands/commandlist.php#_ga=1.116396867.475459535.1470052932

const (
	CmdAcknowledgeHostProblem    = "ACKNOWLEDGE_HOST_PROBLEM"
	CmdAcknowledgeServiceProblem = "ACKNOWLEDGE_SVC_PROBLEM"
	// Schedule downtime and propagate to children
	CmdScheduleRecursiveDowntime = "SCHEDULE_AND_PROPAGATE_HOST_DOWNTIME"

	// Schedule downtime and propagate to childrenon trigger
	CmdScheduleRecursiveTriggeredDowntime = "SCHEDULE_AND_PROPAGATE_TRIGGERED_HOST_DOWNTIME"
	CmdScheduleForcedHostCheck            = "SCHEDULE_FORCED_HOST_CHECK"
	CmdScheduleForcedServiceCheckAll      = "SCHEDULE_FORCED_HOST_SVC_CHECKS"
	CmdScheduleForcedServiceCheck         = "SCHEDULE_FORCED_SVC_CHECK"

	// Schedule downtime of all hosts in hostgroup
	CmdScheduleHostgroupHostsDowntime = "SCHEDULE_HOSTGROUP_HOST_DOWNTIME"

	// Schedule downtime of all services on hosts in hostgroup
	CmdScheduleHostgroupServiceDowntime = "SCHEDULE_HOSTGROUP_SVC_DOWNTIME"
	CmdScheduleHostCheck                = "SCHEDULE_HOST_CHECK"
	CmdScheduleHostDowntime             = "SCHEDULE_HOST_DOWNTIME"
	CmdScheduleHostServiceCheckAll      = "SCHEDULE_HOST_SVC_CHECKS"
	CmdScheduleHostServiceDowntimeAll   = "SCHEDULE_HOST_SVC_DOWNTIME"

	// schedule downtime for all hosts that have service in servicegroup
	CmdScheduleServicegroupHostDowntime = "SCHEDULE_SERVICEGROUP_HOST_DOWNTIME"

	// schedule downtime for all services in the servicegroup
	CmdScheduleServicegroupServiceDowntime = "SCHEDULE_SERVICEGROUP_SVC_DOWNTIME"
	CmdScheduleServiceCheck                = "SCHEDULE_SVC_CHECK"
	CmdScheduleServiceDowntime             = "SCHEDULE_SVC_DOWNTIME"
)

type Command struct {
	Filename string
	cmdFd    io.Writer
}

// Create command interface to a given command file. It should already exist (FIFO created by nagios)
func NewCmd(file string) (c *Command, err error) {
	var cmd Command
	cmd.cmdFd, err = os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	return &cmd, err
}

// Run any nagios command
func (cmd *Command) Cmd(command string, params ...string) (err error) {
	_, err = fmt.Fprintf(cmd.cmdFd, "[%d] %s;%s\n", time.Now().Unix(), command, strings.Join(params, `;`))
	return err
}
