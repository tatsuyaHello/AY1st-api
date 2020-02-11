package handler

import "AY1st/model"

// IsUserSubDupulicateError はUserSubの重複エラーを確認
func IsUserSubDupulicateError(err error) bool {
	errCode, _ := err.(*model.WrappedError)
	return errCode.Code == model.ErrorUserSubDuplicate
}
