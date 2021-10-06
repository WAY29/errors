package errors

import (
	"fmt"
	"io"
)

type Error struct {
	msg           string
	originalError error
	*stack
}

type causer interface {
	Cause() error
}

func New(msg string) error {
	return Error{msg: msg, originalError: nil, stack: callers()}
}

func Newf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return Error{msg: msg, originalError: nil, stack: callers()}
}

func Errorf(format string, args ...interface{}) error {
	return Error{
		msg:           fmt.Sprintf(format, args...),
		originalError: nil,
		stack:         callers(),
	}
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	if msg != "" {
		msg += ": "
	}
	if customErr, ok := err.(Error); ok {
		return Error{
			msg:           fmt.Sprintf("%s%s", msg, customErr.msg),
			originalError: customErr,
			stack:         callers(),
		}
	}

	return Error{msg: fmt.Sprintf("%s%s", msg, err.Error()), originalError: err, stack: callers()}
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, args...)
	if msg != "" {
		msg += ": "
	}

	if customErr, ok := err.(Error); ok {
		return Error{
			msg:           fmt.Sprintf("%s%s", msg, customErr.msg),
			originalError: customErr,
			stack:         callers(),
		}
	}

	return Error{msg: fmt.Sprintf("%s%s", msg, err.Error()), originalError: err, stack: callers()}
}

func Cause(err error) error {
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

func (err Error) Cause() error {
	return err.originalError
}

func (err Error) Error() string {
	return err.msg
}

func (err Error) Unwrap() error {
	return err.originalError
}

func (err Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		io.WriteString(s, err.Error())
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", err.stack)
		}
	case 's':
		io.WriteString(s, err.Error())
	case 'q':
		fmt.Fprintf(s, "%q", err.Error())
	}
}
