# errors

Just another error handling primitives for golang

## Install
```
go install github.com/WAY29/errors@latest
```
## Usage
### New error and print error context
The **errors.New()/errors.Newf()** function returns a new error that has a context, and you can use "%+v" as format descriptor to print error context. For example
```go
err := errors.New("an_error")
fmt.Printf("%+v", err) // will print error message and context information
```

### Print error message
```go
err := errors.New("an_error")
fmt.Printf("%+v", err) // will print error message and context information
fmt.Printf("%#v", err) // will only print error context information
fmt.Printf("%v", err) // will only print error message
fmt.Printf("%s", err) // will only print error message
```
### Adding context to an error
The **errors.Wrap()/errors.Wrapf()** function returns a new error that adds context to the original error. For example
```go
_, err := ioutil.ReadAll(r)
if err != nil {
        return errors.Wrap(err, "read failed")
}
```
### Adding more information to an error
The **errors.Wrap()/errors.Wrapf()** function can also be used to add additional information to wrapped errors. For example
```go
err := errors.New("an_error")
err = errors.Wrap(err, "more information")
fmt.Printf("%+v", err) // will print error message and context information
```

### Retrieving the cause of an error
Using **errors.Wrap()** constructs a stack of errors, adding context to the preceding error. Depending on the nature of the error it may be necessary to reverse the operation of errors.Wrap to retrieve the original error for inspection. Any error value which implements this interface can be inspected by errors.Cause.
```go
type causer interface {
        Cause() error
}
```
errors.Cause will recursively retrieve the topmost error which does not implement causer, which is assumed to be the original cause. For example:
```go
switch err := errors.Cause(err).(type) {
case *MyError:
        // handle specifically
default:
        // unknown error
}
```

### How I use this library
I usually encapsulate this library. For example
```go
package errors

import (
	"fmt"

	"github.com/WAY29/errors"
)

type ErrorType uint16

const (
	Unknown ErrorType = iota
	ProxyError
	RequestError
	ResponseError
)

var (
    DebugFlag = true
)

type CustomError struct {
	Type ErrorType
	Msg  string
}

func (err CustomError) Error() string {
	return err.Msg
}

func New(Type ErrorType, msg string) error {
	return errors.Wrap(CustomError{Type: Type, Msg: msg}, "")
}

func Newf(Type ErrorType, format string, args ...interface{}) error {
	return errors.Wrap(CustomError{Type: Type, Msg: fmt.Sprintf(format, args...)}, "")
}

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// PrintError
func PrintError(err error) {
	// print error context if debug
	if DebugFlag {
		switch customErr := errors.Cause(err).(type) {
		case CustomError:
			switch customErr.Type {
            // case RequestError:
            // case ResponseError:
			// case ProxyError:
			default:
				fmt.Printf("%s: %+v", "Known Error", err)
			}
		default:
			// raw error
            fmt.Printf("%s: %+v", "Raw Error", err)
		}
	} else {
        fmt.Printf("%v", err)
	}
}
```


## Notice
- If you run golang files by `go run`, please run `errors.SetCurrentAbsPath()` first, or stack message about path will be absolute path.
- If you want to skip some frame about stack, please run `errors.SetSkipFrameNum(skipNum)`, this is usually used for your secondary encapsulation of the library.

## Reference
- [pkg/errors](https://github.com/pkg/errors)