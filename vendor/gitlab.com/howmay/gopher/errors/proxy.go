package errors

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

// PkgNew is as the proxy for github.com/pkg/errors.New func.
// func New(message string) error {
// 	return errors.New(message)
// }
var PkgNew = errors.New

// New 包裝 github.com/pkg/errors 的 New
func New(message string, a ...interface{}) error {
	// 用來 handle 過去的 errors.New
	appErrPatten := regexp.MustCompile(`\d{6}`)
	if appErrPatten.MatchString(message) {
		statusCode, err := strconv.Atoi(message[0:3])
		if err != nil {
			statusCode = 400
		}

		newErr := &_error{Code: message, Status: statusCode}
		if len(a) > 0 {
			newErr.Message = a[0].(string)
			// 還有其他參數
			if len(a) > 1 {
				newErr.Message = fmt.Sprintf(newErr.Message, a[1:]...)
			}
		}
		return errors.WithStack(newErr)
	}

	return errors.WithStack(errors.New(message))
}

// Errorf is as the proxy for github.com/pkg/errors.Errorf func.
// func Errorf(format string, args ...interface{}) error {
// 	return errors.Errorf(format, args...)
// }
var Errorf = errors.Errorf

// Wrap is as the proxy for github.com/pkg/errors.Wrap func.
// func Wrap(err error, message string) error {
// 	return errors.Wrap(err, message)
// }
var Wrap = errors.Wrap

// Wrapf is as the proxy for github.com/pkg/errors.Wrapf func.
// func Wrapf(err error, format string, args ...interface{}) error {
// 	return errors.Wrapf(err, format, args...)
// }
var Wrapf = errors.Wrapf

// WithMessage is as the proxy for github.com/pkg/errors.WithMessage func.
// func WithMessage(err error, message string) error {
// 	return errors.WithMessage(err, message)
// }
var WithMessage = errors.WithMessage

// WithMessagef is as the proxy for github.com/pkg/errors.WithMessagef func.
// func WithMessagef(err error, format string, args ...interface{}) error {
// 	return errors.WithMessagef(err, format, args...)
// }
var WithMessagef = errors.WithMessagef

// Cause is as the proxy for github.com/pkg/errors.Cause func.
// func Cause(err error) error {
// 	return errors.Cause(err)
// }
//var Cause = errors.Cause

// WithStack is as the proxy for github.com/pkg/errors.WithStack func.
// func WithStack(err error) error {
// 	return errors.WithStack(err)
// }
var WithStack = errors.WithStack

// Is reports whether any error in err's chain matches target.
// The chain consists of err itself followed by the sequence of errors obtained by repeatedly calling Unwrap.
// An error is considered to match a target if it is equal to that target or if it implements a method Is(error) bool such that Is(target) returns true.
var Is = errors.Is

// Cause ...
var Cause = errors.Cause
