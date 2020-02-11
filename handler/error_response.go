package handler

import "encoding/json"

// ErrorResponse の定義
type ErrorResponse struct {
	Errors []*ErrorResponseInner `json:"errors"`
}

// Append は ErrorResponseにエラーを追加します
func (res *ErrorResponse) Append(code string, t ErrorType, msg string, detail ...interface{}) {
	errRes := &ErrorResponseInner{
		Message: msg,
		Code:    code,
		Type:    t,
		Detail:  detail,
	}
	res.Errors = append(res.Errors, errRes)
}

// String is a stringer impl
func (res ErrorResponse) String() string {
	str, _ := json.MarshalIndent(&res, "", "  ")
	return string(str)
}

// ErrorResponseInner の定義
type ErrorResponseInner struct {
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Type    ErrorType   `json:"type"`
	Detail  interface{} `json:"detail"`
}

// ErrorType エラータイプ
type ErrorType string

const (
	// ErrorAuth 認証エラー
	ErrorAuth ErrorType = "AuthError"
	// ErrorUnknown 不明なエラー
	ErrorUnknown ErrorType = "UnknownError"
	// ErrorParam パラメータエラー
	ErrorParam ErrorType = "ParamError"
	// ErrorNotFound 対象なしエラー
	ErrorNotFound ErrorType = "NotFoundError"
	// ErrorLimitExceeded 制限によるエラー
	ErrorLimitExceeded ErrorType = "LimitExceededError"
)

// NewErrorResponse APIエラー時のレスポンスを生成
func NewErrorResponse(code string, t ErrorType, msg string) ErrorResponse {
	res := ErrorResponse{
		Errors: []*ErrorResponseInner{
			&ErrorResponseInner{
				Message: msg,
				Code:    code,
				Type:    t,
			}}}
	return res
}

// NewErrorResponseDetailed APIエラー時の詳細レスポンスを生成
func NewErrorResponseDetailed(code string, t ErrorType, msg string, detail ...interface{}) ErrorResponse {
	res := ErrorResponse{
		Errors: []*ErrorResponseInner{
			&ErrorResponseInner{
				Message: msg,
				Code:    code,
				Type:    t,
				Detail:  detail,
			}}}
	return res
}
