/* Service definition */
package nagios

type Service struct {
	Hostname string
	HostGroups map[string]string
	Description string
	DisplayName string
	ServiceGroups map[string]string // list of service groups this service belongs to
	Volatile bool
	Contacts map[string]string // list of service contacts
	ContactGroups map[string]string // list of service contact groups
}


func NewService() Service {
	var s Service
	s.HostGroups = make(map[string]string)
	s.ServiceGroups = make(map[string]string)
	s.Contacts = make(map[string]string)
	s.ContactGroups = make(map[string]string)
	return s
}
