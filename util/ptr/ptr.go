package ptr

import "time"

// Uint8 はuint8型の値を受け取り、そのポインタを返します。
func Uint8(u uint8) *uint8 {
	return &u
}

// Uint16 はuint16型の値を受け取り、そのポインタを返します。
func Uint16(u uint16) *uint16 {
	return &u
}

// Uint32 はuint32型の値を受け取り、そのポインタを返します。
func Uint32(u uint32) *uint32 {
	return &u
}

// Uint64 はuint64型の値を受け取り、そのポインタを返します。
func Uint64(u uint64) *uint64 {
	return &u
}

// Int8 はint8型の値を受け取り、そのポインタを返します。
func Int8(i int8) *int8 {
	return &i
}

// Int16 はint16型の値を受け取り、そのポインタを返します。
func Int16(i int16) *int16 {
	return &i
}

// Int32 はint32型の値を受け取り、そのポインタを返します。
func Int32(i int32) *int32 {
	return &i
}

// Int64 はint64型の値を受け取り、そのポインタを返します。
func Int64(i int64) *int64 {
	return &i
}

// Float32 はfloat32型の値を受け取り、そのポインタを返します。
func Float32(f float32) *float32 {
	return &f
}

// Float64 はfloat64型の値を受け取り、そのポインタを返します。
func Float64(f float64) *float64 {
	return &f
}

// Complex64 はcomplex64型の値を受け取り、そのポインタを返します。
func Complex64(c complex64) *complex64 {
	return &c
}

// Complex128 はcomplex128型の値を受け取り、そのポインタを返します。
func Complex128(c complex128) *complex128 {
	return &c
}

// Byte はbyte型の値を受け取り、そのポインタを返します。
func Byte(b byte) *byte {
	return &b
}

// Rune はrune型の値を受け取り、そのポインタを返します。
func Rune(r rune) *rune {
	return &r
}

// Uint はuint型の値を受け取り、そのポインタを返します。
func Uint(u uint) *uint {
	return &u
}

// Int はint型の値を受け取り、そのポインタを返します。
func Int(i int) *int {
	return &i
}

// String はstring型の値を受け取り、そのポインタを返します。
func String(s string) *string {
	return &s
}

// Bool はbool型の値を受け取り、そのポインタを返します。
func Bool(b bool) *bool {
	return &b
}

// True はtrue(bool型)のポインタを返します。
func True() *bool {
	return Bool(true)
}

// False はfalse(bool型)のポインタを返します。
func False() *bool {
	return Bool(false)
}

// Time はtime.Time型の値を受け取り、そのポインタを返します。
func Time(t time.Time) *time.Time {
	return &t
}
