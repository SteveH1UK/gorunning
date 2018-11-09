package validation

import (
	"strconv"
	"strings"
	"time"

	"github.com/SteveH1UK/gorunning"
	"github.com/SteveH1UK/gorunning/utils"
)

const minAge = 16

// ValidateAthelete - Validates a new or Updated Athelete
func ValidateAthelete(a root.NewAthelete) []root.ValidationError {
	var validationErrors []root.ValidationError
	if len(strings.TrimSpace(a.FriendyURL)) < 4 {
		validationError := root.ValidationError{Code: 701, Field: "friendly_url", Message: "must be at least 4 characters long"}
		validationErrors = append(validationErrors, validationError)
	}

	if !utils.ContainsAllowedCharactersOnly(a.FriendyURL) {
		validationError := root.ValidationError{Code: 702, Field: "friendly_url", Message: "must only contain alphanumerics or hypen or underscore"}
		validationErrors = append(validationErrors, validationError)
	}

	if !utils.ValidateDate(a.DateOfBirth) {
		validationError := root.ValidationError{Code: 703, Field: "date-of-birth", Message: "must be in format dd-MM-yyyy"}
		validationErrors = append(validationErrors, validationError)
	}

	// TODO - check age is at least a certain age
	if !IsOldEnough(a.DateOfBirth, minAge) {
		validationError := root.ValidationError{Code: 704, Field: "date-of-birth", Message: "Must be at least " + strconv.Itoa(minAge)}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors
}

// IsOldEnough assumes date is correct format
func IsOldEnough(dateOfBirth string, minYears int) bool {
	dob, _ := utils.ParseApplicationDate(dateOfBirth)

	age := time.Since(dob)
	years := age.Hours() / (24 * 365.25)

	if int(years)-minYears > 0 {
		return true
	}
	return false
}
