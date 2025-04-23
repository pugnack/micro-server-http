package http

import "fmt"

// SetError pass error to caller
func SetError(err interface{}) error {
	return &Error{err: err}
}

// GetError return underline error
func GetError(err interface{}) interface{} {
	if verr, ok := err.(*Error); ok {
		return verr.err
	}
	return err
}

// Error struct holds error
type Error struct {
	err interface{}
}

// Error func for error interface
func (err *Error) Error() string {
	return fmt.Sprintf("%v", err.err)
}
