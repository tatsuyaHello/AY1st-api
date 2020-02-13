package model

// ErrorCode エラーコード
type ErrorCode string

func (ec ErrorCode) Error() string {
	return string(ec)
}

const (
	// ErrorRequiredQueryParameterMissing 必須のクエリパラメータのkey-value両方が無い
	ErrorRequiredQueryParameterMissing ErrorCode = "RequiredQueryParameterMissing"
	// ErrorInvalidQueryParameterValue クエリパラメータに不正な値が渡された
	ErrorInvalidQueryParameterValue ErrorCode = "InvalidQueryParameterValue"
	// ErrorRequestBodyMismatch リクエストデータの形式が異なります
	ErrorRequestBodyMismatch ErrorCode = "RequestBodyMismatch"
	// ErrorResourceNotFound 条件に一致するデータが見つかりません
	ErrorResourceNotFound ErrorCode = "ResourceNotFound"
	// ErrorConflict リソースの現在の状態と矛盾する操作を行おうとした
	ErrorConflict ErrorCode = "Conflict"
	// ErrorNotAcceptable 残りの空きリソースがないなどの事情でリクエストを受け入れられない
	ErrorNotAcceptable ErrorCode = "NotAcceptable"
	// ErrorUserSubDuplicate ユーザサブの重複
	ErrorUserSubDuplicate ErrorCode = "IdentityCodeDuplicate"
	// ErrorUnauthorized ユーザーが存在しないなど、トークン由来以外の認証エラー
	ErrorUnauthorized ErrorCode = "Unauthorized"
	// ErrorInvalidToken 認証トークンが無効
	ErrorInvalidToken ErrorCode = "InvalidToken"
	// ErrorNoAccessAuthority 対象データにアクセスする権限がありません
	ErrorNoAccessAuthority ErrorCode = "NoAccessAuthority"
	// ErrorInternalServerError 原因不明の障害
	ErrorInternalServerError ErrorCode = "InternalServerError"
	// ErrorUnderMaintenance メンテナンス中
	ErrorUnderMaintenance ErrorCode = "UnderMaintenance"
	// ErrorCannotCreate 作成に失敗
	ErrorCannotCreate ErrorCode = "CannotCreate"
)

// WrappedError は外向きのエラー
type WrappedError struct {
	Code    ErrorCode
	Message string
	cause   error
	Detail  []interface{}
}

func (e *WrappedError) Error() string {
	return e.Code.Error()
}

// Cause は元のエラーを取り出す
func (e *WrappedError) Cause() error {
	return e.cause
}

// WrapError returns エラーコードを指定してラップしたエラーを生成します。
func WrapError(cause error, code ErrorCode, msg string) error {
	we := &WrappedError{Code: code, Message: msg, cause: cause}
	return we
}

// NewError returns エラーコードを指定してエラーを生成します。
func NewError(code ErrorCode, msg string) error {
	we := &WrappedError{Code: code, Message: msg}
	return we
}
