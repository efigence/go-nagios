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


func NewNotificationFromEnv() Notification {
    n := NewNotification()
	// FIXME handle error
	n.Host, _ = NewHostFromEnv()
	n.Service, _ = NewServiceFromEnv()
	return n
}
