package errcode

import (
	"fmt"
	"net/http"
)

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

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case UnauthorizedAuth.Code:
		fallthrough
	case UnauthorizedToken.Code:
		fallthrough
	case TokenTimeout.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		return http.StatusUnauthorized
	case TooManyRequest.Code:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}
