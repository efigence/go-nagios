/* Host definition */
package nagios

type Host struct {
	Hostname string
	Displayname string
	HostGroups map[string]string
	Address string
	Parents *[]Host
}
