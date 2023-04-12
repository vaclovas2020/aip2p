// Package implements AIP2P node logging functionality
package logs

import "log"

// Log success message with vendor prefix
func Log(message string, vendor string) {
	log.Printf("\033[32m[%s][success]\033[0m %s", vendor, message)
}

// Log error message with vendor prefix
func LogError(e error, vendor string, callPanic bool) {
	log.Printf("\033[31m[%s][error]\033[0m %s", vendor, e.Error())
	if callPanic {
		panic(e)
	}
}

// Log warning message with vendor prefix
func LogWarning(message string, vendor string) {
	log.Printf("\033[33m[%s][warning]\033[0m %s", vendor, message)
}
