package errcode

import "fmt"

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

var codes = map[int]string{}

func newError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic("error code is used!")
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = []string{}
	for _, d := range details {
		newError.Details = append(newError.Details, d)
	}
	return &newError
}
