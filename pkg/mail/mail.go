package mail

import (
	"gohub/pkg/config"
	"sync"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte // Planintext message (optional)
	HTML    []byte // HTML message (optional)
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

// NewMailer singleton
func NewMailer() *Mailer {
	once.Do(func() {
		internalMailer = &Mailer{
			Driver: &SMTP{},
		}
	})
	return internalMailer
}

// Send implementation send email
func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, config.GetStringMapString("mail.smtp"))
}
