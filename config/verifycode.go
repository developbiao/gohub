package config

import "gohub/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{
			// Verify code length
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),
			// Expire time
			"expire_time": config.Env("VERIFY_CODE_EXPIRE", 15),
			// Debug expire time
			"debug_expire_time": 10080,
			// Debug default verify code
			"debug_code": 123456,
			// Debug phone default prefix
			"debug_phone_prefix": "000",
			// Debug testing email suffix
			"debug_email_suffix": "@testing.com",
		}
	})
}
