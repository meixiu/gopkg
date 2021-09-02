package errcode

import (
	"fmt"
)

var errList = make(map[int]Error)

type Error interface {
	Error() string
	Code() int
	Message() string
}

type codeError struct {
	code    int
	message string
}

func (ce *codeError) Error() string {
	return fmt.Sprintf("(code=%d) %s", ce.code, ce.message)
}

func (ce *codeError) Code() int {
	return ce.code
}

func (ce *codeError) Message() string {
	return ce.message
}

func New(code int, msg string) Error {
	return add(code, msg)
}

func Is(err error) bool {
	_, ok := err.(Error)
	return ok
}

func As(err error) (Error, bool) {
	if Is(err) {
		return err.(Error), true
	}
	return nil, false
}

func add(code int, msg string) Error {
	if _, has := errList[code]; has {
		panic(fmt.Sprintf("ecode: %d already exist", code))
	}
	ce := &codeError{
		code:    code,
		message: msg,
	}
	errList[code] = ce
	return ce
}

func CodeList() map[int]string {
	l := make(map[int]string)
	for k, v := range errList {
		l[k] = v.Message()
	}
	return l
}
