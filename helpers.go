package configuration

import (
	"log"
)

var (
	gLoggingEnabled  bool
	gFailIfCannotSet bool
)

func logf(format string, args ...interface{}) {
	if gLoggingEnabled {
		log.Printf(format, args...)
	}
}

func fail(format string, args ...interface{}) {
	if gFailIfCannotSet {
		log.Fatalf(format, args...)
	}
}
