package sms

// Driver sms interface must implement Send method
type Driver interface {
	Send(phone string, message Message, config map[string]string) bool
}
