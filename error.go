package envparser

import (
	"errors"
	"fmt"
)

var (
	ErrEnvVarEmpty    = errors.New("environment variable empty")
	ErrInvalidVarName = errors.New("invalid variable name")
)

type ErrorWrapper struct {
	KeyName string
	Err     error
}

func (e *ErrorWrapper) Error() string {
	return fmt.Sprintf(`env variable name \"%s\" error: %s`, e.KeyName, e.Err)
}

func (e *ErrorWrapper) Unwrap() error {
	return e.Err
}

// ErrorIsEnvVarEmpty return true if error is same with ErrEnvVarEmpty
func ErrorIsEnvVarEmpty(err error) bool {
	return errors.Is(err, ErrEnvVarEmpty)
}

// ErrorIsInvalidVarName return true if error is same with ErrInvalidVarName
func ErrorIsInvalidVarName(err error) bool {
	return errors.Is(err, ErrInvalidVarName)
}
