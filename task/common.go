package task

import "fmt"

type StatusCodeError struct {
	code int
}

func (r StatusCodeError) Error() string {
	return fmt.Sprintf("status code is %d", r.code)
}

func FormatStatusCodeError(status int) StatusCodeError {
	return StatusCodeError{
		code: status,
	}
}
