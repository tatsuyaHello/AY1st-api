package model

import (
	"encoding/json"
	"fmt"
)

// ErrorResponse の定義
type ErrorResponse struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// String is a stringer impl
func (res ErrorResponse) String() string {
	str, _ := json.MarshalIndent(&res, "", "  ")
	return string(str)
}

// NewErrorResponse APIエラー時のレスポンスを、 model.WrappedError から生成
func NewErrorResponse(err error) *ErrorResponse {
	// ここのType Assertion は 失敗すればpanicになるが、実装漏れが原因なのでハンドリングしない。
	wrappedErr, ok := err.(*WrappedError)
	if !ok {
		panic(fmt.Errorf("error is not *model.WrappedError, type=%T, error = %+v", err, err))
	}
	res := &ErrorResponse{
		Code:    wrappedErr.Code,
		Message: wrappedErr.Message,
	}
	return res
}

// NewErrorResponseWithCode APIエラー時のレスポンスをコードを指定して生成
func NewErrorResponseWithCode(code ErrorCode, message interface{}) *ErrorResponse {
	res := &ErrorResponse{
		Code:    code,
		Message: msgToString(message),
	}
	return res
}

func msgToString(message interface{}) string {
	switch m := message.(type) {
	case string:
		return m
	case error:
		return m.Error()
	case fmt.Stringer:
		return m.String()
	default:
		return fmt.Sprint(m)
	}
}
