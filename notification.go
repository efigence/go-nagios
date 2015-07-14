// Nagios notification manipulation
// see doc/notification.md for details about input format

package nagios

type Notification struct {
	Host Host
	Service Service
}


func NewNotification() Notification {
	var n Notification
	n.Host = NewHost()
	n.Service = NewService()
	return n
}


func NewNotificationFromEnv() (Notification, error) {
    n := NewNotification()
	// FIXME handle error
	var err error
	n.Host, err = NewHostFromEnv()
	n.Service, err = NewServiceFromEnv()
	return n, err
}
