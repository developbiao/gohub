package mail

import (
	"fmt"
	emailPKG "github.com/jordan-wright/email"
	"gohub/pkg/logger"
	"net/smtp"
)

// SMTP implementation email.Driver interface
type SMTP struct{}

// Send implementation send email
func (s *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("Send mail", "Send detail", e)

	err := e.Send(fmt.Sprintf("%v:%v", config["host"], config["port"]), smtp.PlainAuth(
		"",
		config["username"],
		config["password"],
		config["host"],
	))
	if err != nil {
		logger.ErrorString("Send mail", "Send mail error", err.Error())
		return false
	}
	logger.DebugString("Send mail", "Send mail success", "")
	return true
}
