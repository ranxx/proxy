package errors

import (
	"fmt"
)

// Response ...
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ErrCode ...
type ErrCode struct {
	Code ErrNumber `json:"code"`
	Msg  string    `json:"msg"`
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
	case *ErrCode:
		ecode := err.(*ErrCode)
		resp.Code = int(ecode.Code)
		resp.Msg = ecode.Msg
	default:
		resp.Code = int(Err)
		resp.Msg = err.Error()
	}
	return resp
}

// NewNotFound msg
func NewNotFound(msg string) *ErrCode {
	return newErrorCode(ErrNotFound, msg)
}

// NewErrCode ...
func NewErrCode(code ErrNumber, msg string) *ErrCode {
	return newErrorCode(code, msg)
}

func newErrorCode(code ErrNumber, msg string) *ErrCode {
	return &ErrCode{Code: code, Msg: msg}
}
