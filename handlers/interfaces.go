package handlers

// InfoLogger logs errors.
type ErrorLogger interface {
	Errorf(format string, args ...interface{})
}
