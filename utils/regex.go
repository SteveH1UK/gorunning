package utils

import (
	"regexp"
)

// ContainsAllowedCharactersOnly - contains only Alpha, numbers, underscore and hyphen
func ContainsAllowedCharactersOnly(input string) bool {

	re := regexp.MustCompile("^[a-zA-Z0-9_-]*$")
	return re.MatchString(input)
}
