package sms

import (
	"gohub/pkg/config"
	"sync"
)

// Message struct
type Message struct {
	Template string
	Data     map[string]string

	Content string
}

// SMS operator driver
type SMS struct {
	Driver Driver
}

// once single instance
var once sync.Once

// internalSMS internal SMS object
var internalSMS *SMS

// NewSMS new sms
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return internalSMS
}

// Send implement send
func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
