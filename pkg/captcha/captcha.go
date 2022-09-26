package captcha

import (
	"github.com/mojocn/base64Captcha"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"sync"
)

// Captcha instance
type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// once ensure internalCaptcha is only init one
var once sync.Once

// internalCaptcha for internal using
var internalCaptcha *Captcha

// NewCaptcha new captcha instance
func NewCaptcha() *Captcha {
	once.Do(func() {
		// Initialization Captcha object
		internalCaptcha = &Captcha{}

		// Global use Redis object, and config prefix key
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha",
		}

		// Config base64Captcha driver
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),      // height
			config.GetInt("captcha.width"),       // width
			config.GetInt("captcha.length"),      // length
			config.GetFloat64("captcha.maxskew"), // max absolute skew factor of a single digit
			config.GetInt("captcha.dotcount"),    // Number of background circles.
		)
		// New base64 captcha instance
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})
	return internalCaptcha
}

// GenerateCaptcha implement generate captcha
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha implement verify captcha
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}
	return c.Base64Captcha.Verify(id, answer, false)
}
