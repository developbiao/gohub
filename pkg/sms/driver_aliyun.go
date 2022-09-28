package sms

import (
	"encoding/json"
	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
	"gohub/pkg/logger"
)

// Aliyun implement sms.Driver interface
type Aliyun struct{}

func (a *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	smsClient := aliyunsmsclient.New("http://dysmsapi.aliyuncs.com/")

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("SMS [aliyun]", "Paring bind error", err.Error())
	}
	logger.DebugJSON("SMS [aliyun]", "config info", config)

	result, err := smsClient.Execute(
		config["access_key_id"],
		config["access_key_secret"],
		phone,
		config["sign_name"],
		message.Template,
		string(templateParam),
	)
	logger.DebugJSON("SMS [aliyun]", "Request", smsClient.Request)
	logger.DebugJSON("SMS [aliyun]", "Result", result)

	if err != nil {
		logger.ErrorString("SMS [aliyun]", "send failed", err.Error())
		return false
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		logger.ErrorString("SMS [aliyun]", "parsing failed", err.Error())
		return false
	}
	if result.IsSuccessful() {
		logger.ErrorString("SMS [aliyun]", "send success", string(resultJSON))
		return true
	} else {
		logger.ErrorString("SMS [aliyun]", "sms provider response error", string(resultJSON))
		return false
	}
}
