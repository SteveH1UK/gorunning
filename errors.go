package root

import "errors"

// ValidationError - one validation error
type ValidationError struct {
	Code    int    `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error - application error
type Error struct {
	HTTPCode         int               `json:"http-code"`
	Message          string            `json:"message"`
	ValidationErrors []ValidationError `json:"validation-errors,omitempty"`
}

var (
	// ErrDBErrNotFound error returned when a document could not be found
	ErrDBErrNotFound = errors.New("not found on Database")
	// ErrDBRecordExists - record already present on the database
	ErrDBRecordExists = errors.New("Record Exists on Database")
)
