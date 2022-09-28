package config

import "gohub/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{
			// Default aliyun sign_name and template_code
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_ALIYUN_SIGN_NAME", "阿里云短信测试"),
				"template_code":     config.Env("SMS_ALIYUN_TEMPLATE_CODE", "SMS_123456"),
			},
			// For test without third part
			"test": map[string]interface{}{
				"app_id":        config.Env("SMS_TEST_APP_ID"),
				"app_key":       config.Env("SMS_TEST_APP_KEY"),
				"template_code": config.Env("SMS_TEST_TEMPLATE_CODE", "SMS_123456"),
			},
		}
	})
}
