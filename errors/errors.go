package errors

import (
	"fmt"
	"net/http"
)

// Response ...
type Response struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	HTTPStatus int         `json:"-"`
}

// ErrCode ...
type ErrCode struct {
	Code       ErrNumber `json:"code"`
	Msg        string    `json:"msg"`
	HTTPStatus int       `json:"-"`
}

func (e ErrCode) Error() string {
	return fmt.Sprintf("code:%d msg:%s", e.Code, e.Msg)
}

// Convert2Response ...
func Convert2Response(err error) Response {
	resp := Response{Data: nil}
	switch err.(type) {
	case ErrCode:
		ecode := err.(ErrCode)
		resp.Code = int(ecode.Code)
		resp.Msg = ecode.Msg
		resp.HTTPStatus = ecode.HTTPStatus
	case *ErrCode:
		ecode := err.(*ErrCode)
		resp.Code = int(ecode.Code)
		resp.Msg = ecode.Msg
		resp.HTTPStatus = ecode.HTTPStatus
	default:
		resp.Code = int(Err)
		resp.Msg = err.Error()
		resp.HTTPStatus = 200
	}
	return resp
}

// NewNotFound msg
func NewNotFound(format string, a ...interface{}) *ErrCode {
	msg := fmt.Sprintf(format, a...)
	return newErrorCode(ErrNotFound, msg)
}

// NewHTTPErr ...
// status default 200
func NewHTTPErr(code ErrNumber, msg string, status ...int) *ErrCode {
	st := 200
	if len(status) > 0 {
		st = status[0]
	}
	return &ErrCode{
		Code:       code,
		Msg:        msg,
		HTTPStatus: st,
	}
}

// NewHTTPBadRequest 400
func NewHTTPBadRequest(msg ...string) *ErrCode {
	str := "请求参数解析失败"
	if len(msg) > 0 {
		str = msg[0]
	}
	return &ErrCode{
		Code:       http.StatusBadRequest,
		Msg:        str,
		HTTPStatus: http.StatusBadRequest,
	}
}

// NewErrCode ...
func NewErrCode(code ErrNumber, msg string) *ErrCode {
	return newErrorCode(code, msg)
}

func newErrorCode(code ErrNumber, msg string) *ErrCode {
	return &ErrCode{Code: code, Msg: msg}
}
