package exception

import "errors"

var (
	// ErrValidationFailed is returned when input validation fails
	ErrValidationFailed = errors.New("validation failed")

	// ErrServiceUnavailable is returned when a required service is unavailable
	ErrServiceUnavailable = errors.New("service unavailable")
)
