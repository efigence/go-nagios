package nagios
type Summary struct {
	HostCount `json:"all_host"`
	ServiceCount `json:"all_service"`
}

type HostCount struct {
	All int `json:"all"`
	Up int `json:"up"`
	Down int `json:"down"`
	Unreachable int `json:"unreachable"`
	Downtime int `json:"downtime"`
	Acknowledged int `json:"ack"`
}

type ServiceCount struct {
	All int `json:"all"`
	Ok int `json:"up"`
	Warning int `json:"down"`
	Critical int `json:"unreachable"`
	Unknown int `json:"unreachable"`
	Downtime int `json:"downtime"`
	Acknowledged int `json:"ack"`
}

func (sum *Summary) UpdateHost(hosts map[string]Host) error {
	var err error
	var c HostCount
	for _, v := range hosts {
		c.All++
		if v.Acknowledged {
			c.Acknowledged++
		}
		if v.Downtime {
			c.Downtime++
		}
		if v.State == StateUp {
			c.Up++
		}
		if v.State == StateDown {
			c.Down++
		}
		if v.State == StateUnreachable {
			c.Unreachable++
		}
	}
	sum.HostCount = c
	return err
}

func (sum *Summary) UpdateService(services map[string]map[string]Service) error {
	var err error
	var c ServiceCount
	for _, v1 := range services {
		for _, v2 := range v1 {
			c.All++
			if v2.Acknowledged {
				c.Acknowledged++
			}
			if v2.Downtime {
				c.Downtime++
			}
			if v2.State == StateOk {
				c.Ok++
			}
			if v2.State == StateWarning {
				c.Warning++
			}
			if v2.State == StateCritical {
				c.Critical++
			}
			if v2.State == StateUnknown {
				c.Unknown++
			}
		}
	}
	sum.ServiceCount = c
	return err
}
