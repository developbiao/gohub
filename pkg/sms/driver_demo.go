package sms

import (
	"gohub/pkg/logger"
	"regexp"
)

// Demo sms driver implement sms.Driver
type Demo struct{}

func (d *Demo) Send(phone string, message Message, config map[string]string) bool {
	re := regexp.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$")
	if re.MatchString(phone) {
		logger.DebugString("SMS [TEST]", "send success", message.Content)
		return true
	} else {
		logger.ErrorString("SMS [TEST]", "send failed", message.Content)
		return false
	}
}
