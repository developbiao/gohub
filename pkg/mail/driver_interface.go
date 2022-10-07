package mail

// Driver email driver interface
type Driver interface {
	// Send email check captcha answer code
	Send(email Email, config map[string]string) bool
}
