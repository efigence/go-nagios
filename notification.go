// Nagios notification manipulation
// see doc/notification.md for details about input format

package nagios

type Notification struct {
	Host Host
	Service Service
}


func NewNotification() Notification {
	var n Notification
	n.Host = NewHostFromEnv()
	n.Service = NewServiceFromEnv()
	return n
}


func NewNotificationFromEnv() Notification {
	var n Notification
	n.Host = NewHostFromEnv()
	n.Service = NewServiceFromEnv()
	return n
}
