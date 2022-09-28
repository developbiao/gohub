package verifycode

type Store interface {
	// Set Save verify code
	Set(id string, value string) bool

	// Get verify code
	Get(id string, clear bool) string

	// Verify verify code
	Verify(id, answer string, clear bool) bool
}
