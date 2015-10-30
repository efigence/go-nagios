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
	return err
}

func (sum *Summary) UpdateService(map[string]map[string]Service) error {
	var err error
	return err
}
