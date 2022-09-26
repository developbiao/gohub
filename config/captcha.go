package config

import "gohub/pkg/config"

func init() {
	config.Add("captcha", func() map[string]interface{} {
		return map[string]interface{}{
			"height":            80,                  // height
			"width":             240,                 // width
			"length":            6,                   // length
			"maxske":            0.7,                 // max absolute skew factor of a single digit
			"dotcount":          80,                  // Number of background circles.
			"expire_time":       15,                  // Expire time 15 minutes
			"debug_expire_time": 10080,               // For debug expire time
			"testing_key":       "captcha_skip_test", // For debug testing key
		}
	})
}
