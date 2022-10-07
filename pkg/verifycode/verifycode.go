package verifycode

import (
	"fmt"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/helpers"
	"gohub/pkg/logger"
	"gohub/pkg/mail"
	"gohub/pkg/redis"
	"gohub/pkg/sms"
	"strings"
	"sync"
)

// VerifyCode struct
type VerifyCode struct {
	Store Store
}

// once singleton instance
var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				// Set key prefix keep tidy
				KeyPrefix: config.GetString("app.name") + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

func (vc *VerifyCode) SendSMS(phone string) bool {
	// Generate verify code
	code := vc.generateVerifyCode(phone)
	// Easement for local debug
	if !app.IsProduction() &&
		strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.test.template_code"),
		Data:     map[string]string{"code": code},
	})
}

// generateVerifyCode generate random verify code
func (vc *VerifyCode) generateVerifyCode(key string) string {
	var code string
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	} else {
		code = helpers.RandomNumber(config.GetInt("verifycode.code_length"))
	}
	// Set verify code to cache
	vc.Store.Set(key, code)
	return code
}

// CheckAnswer verify code is correct
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	logger.DebugJSON("verifycode", "CheckAnswer", map[string]string{key: answer})
	if !app.IsProduction() &&
		strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	return vc.Store.Verify(key, answer, false)
}

// SendEmail send email captcha code
func (vc *VerifyCode) SendEmail(email string) error {
	// Generate verify code
	code := vc.generateVerifyCode(email)

	// Easement for local debug
	if !app.IsProduction() &&
		strings.HasSuffix(email, config.GetString("verifycode.debug_email_suffix")) {
		logger.DebugString("send email", "SUFFIX DEBUG", config.GetString("verifycode.debug_email_suffix"))
		return nil
	}

	content := fmt.Sprintf("<h3>Your Email Captcha code is %v </h3>", code)
	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{email},
		Subject: "Email Captcha code",
		HTML:    []byte(content),
	})
	return nil
}
