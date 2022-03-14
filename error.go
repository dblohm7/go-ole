package ole

import (
	"strings"
)

// OleError stores COM errors.
type OleError struct {
	hr          uintptr
	description string
	subError    error
}

// NewError creates new error with HResult.
func NewError(hr uintptr) *OleError {
	return &OleError{hr: hr}
}

// NewErrorWithDescription creates new COM error with HResult and description.
func NewErrorWithDescription(hr uintptr, description string) *OleError {
	return &OleError{hr: hr, description: description}
}

// NewErrorWithSubError creates new COM error with parent error.
func NewErrorWithSubError(hr uintptr, err error) *OleError {
	return &OleError{hr: hr, subError: err}
}

// Code is the HResult.
func (v *OleError) Code() uintptr {
	return uintptr(v.hr)
}

// String description, either manually set or format message with error code.
func (v *OleError) String() string {
	var sb strings.Builder
	sb.WriteString(errstr(int(v.hr)))
	if v.description != "" {
		sb.WriteString(" (")
		sb.WriteString(v.description)
		sb.WriteString(")")
	}
	if v.subError != nil {
		sb.WriteString(" (")
		sb.WriteString(v.subError.Error())
		sb.WriteString(")")
	}
	return sb.String()
}

// Error implements error interface.
func (v *OleError) Error() string {
	return v.String()
}

// Description retrieves error summary, if there is one.
func (v *OleError) Description() string {
	return v.description
}

// SubError returns parent error, if there is one.
func (v *OleError) SubError() error {
	return v.subError
}

// Unwrap enables OleError to be compatible with the functions in the standard library's errors package
func (v *OleError) Unwrap() error {
	return v.subError
}
