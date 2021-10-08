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
### Adding error type to an error
The **errors.SetType()** function returns a new error that adds error type to wrapped error. For example
```go
type ErrorType uint16

const (
	RequestError ErrorType = iota
	ResponseError
)


func test() {
	err := errors.New("new error")
	err, _ = errors.SetType(err, RequestError)
	// or errors.SetTypeWithoutBool
	// err = errors.SetTypeWithoutBool(err, RequestError)

	switch errType := errors.GetType(err); errType {
	case RequestError:
		fmt.Printf("Request error: %+v\n", err)
	case ResponseError:
		fmt.Printf("Response error: %+v\n", err)
	default:
		fmt.Printf("Unknown error: %#v\n", err)
	}
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
For example
```go
package errors

import (
	"fmt"
	"github.com/WAY29/errors"
)

type ErrorType uint16

const (
	RequestError ErrorType = iota
	ResponseError
)

func main() {
	// errors.SetCurrentAbsPath()

	err := errors.New("new error")
	err, _ = errors.SetType(err, RequestError)
	err = errors.Wrapf(err, "wrapped")

	switch errType := errors.GetType(err); errType {
	case RequestError:
		fmt.Printf("Request error: %+v\n", err)
	case ResponseError:
		fmt.Printf("Response error: %+v\n", err)
	default:
		fmt.Printf("Unknown error: %#v\n", err)
	}
}

```


## Notice
- If you run golang files by `go run`, please run `errors.SetCurrentAbsPath()` first, or stack message about path will be absolute path.
- If you want to skip some frame about stack, please run `errors.SetSkipFrameNum(skipNum)`, this is usually used for your secondary encapsulation of the library.

## Reference
- [pkg/errors](https://github.com/pkg/errors)