package validator

import "gohub/pkg/captcha"

// ValidateCaptcha validation captcha
func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "Captcha answer is error")
	}
	return errs
}
