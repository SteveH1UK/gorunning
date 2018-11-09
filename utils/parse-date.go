package utils

import (
	"time"
)

// ValidateDate in dd-MM-yyyy format
func ValidateDate(value string) bool {

	_, err := ParseApplicationDate(value)
	if err == nil {
		return true
	}
	return false
}

// ParseApplicationDate - parse date in application layout
func ParseApplicationDate(value string) (time.Time, error) {

	layout := "02-01-2006" // magical reference date Mon Jan 2 15:04:05 MST 2006
	return time.Parse(layout, value)
}
