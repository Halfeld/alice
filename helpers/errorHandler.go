package helpers

import "log"

// ErrorHandler abstract error messages
func ErrorHandler(err error, message string) {
	if err != nil {
		log.Fatalf(message, ": %v", err)
	}
}
