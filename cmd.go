package nagios

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// https://old.nagios.org/developerinfo/externalcommands/commandlist.php#_ga=1.116396867.475459535.1470052932
// Schedule downtime and propagate to children
const CmdScheduleRecursiveDowntime = "SCHEDULE_AND_PROPAGATE_HOST_DOWNTIME"

// Schedule downtime and propagate to childrenon trigger
const CmdScheduleRecursiveTriggeredDowntime = "SCHEDULE_AND_PROPAGATE_TRIGGERED_HOST_DOWNTIME"
const CmdScheduleForcedHostCheck = "SCHEDULE_FORCED_HOST_CHECK"
const CmdScheduleForcedServiceCheckAll = "SCHEDULE_FORCED_HOST_SVC_CHECKS"
const CmdScheduleForcedServiceCheck = "SCHEDULE_FORCED_SVC_CHECK"

// Schedule downtime of all hosts in hostgroup
const CmdScheduleHostgroupHostsDowntime = "SCHEDULE_HOSTGROUP_HOST_DOWNTIME"

// Schedule downtime of all services on hosts in hostgroup
const CmdScheduleHostgroupServiceDowntime = "SCHEDULE_HOSTGROUP_SVC_DOWNTIME"
const CmdScheduleHostCheck = "SCHEDULE_HOST_CHECK"
const CmdScheduleHostDowntime = "SCHEDULE_HOST_DOWNTIME"
const CmdScheduleHostServiceCheckAll = "SCHEDULE_HOST_SVC_CHECKS"
const CmdScheduleHostServiceDowntimeAll = "SCHEDULE_HOST_SVC_DOWNTIME"

// schedule downtime for all hosts that have service in servicegroup
const CmdScheduleServicegroupHostDowntime = "SCHEDULE_SERVICEGROUP_HOST_DOWNTIME"

// schedule downtime for all services in the servicegroup
const CmdScheduleServicegroupServiceDowntime = "SCHEDULE_SERVICEGROUP_SVC_DOWNTIME"
const CmdScheduleServiceCheck = "SCHEDULE_SVC_CHECK"
const CmdScheduleServiceDowntime = "SCHEDULE_SVC_DOWNTIME"

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
	_, err = fmt.Fprintf(cmd.cmdFd, "[%d] %s %+v\n", time.Now().Unix(), command, strings.Join(params, `;`))
	return err
}
