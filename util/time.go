package util

import (
	"strings"
	"time"
)

// GetTimeNowFunc is current time function.
// Default function is `time.Now`.
var GetTimeNowFunc = time.Now

// TokyoTimeLocation is Asia/Tokyo time location
var TokyoTimeLocation = time.FixedZone("Asia/Tokyo", 9*60*60)

// RFC3339NanoTrailingZeros `time.RFC3339Nano` を部分的に変更した末尾ゼロを取り除かないフォーマット
const RFC3339NanoTrailingZeros = "2006-01-02T15:04:05.000000000Z07:00"

// DBTimeFormat DB用のフォーマット
const DBTimeFormat = "2006-01-02 15:04:05"

// GetTimeNow は現在時刻を返す。
// `GetTimeNowFunc` を上書きすることで、テストなどの用途で任意の日時を取得するように調整できる。
func GetTimeNow() time.Time {
	now := GetTimeNowFunc().In(TokyoTimeLocation)
	return now
}

// GetFormatedTimeNow はDB用フォーマット済みの現在時刻を返す
func GetFormatedTimeNow() string {
	return GetTimeNow().Format(DBTimeFormat)
}

// EqualTime 文字列から日付を比較
// '/' と '-' 区切りの日付のみ考慮
func EqualTime(tA, tB string) bool {

	timeA := strings.ReplaceAll(tA, "/", "-")

	FormatedTimeA, err := time.Parse(DBTimeFormat, timeA)

	if err != nil {
		return false
	}

	timeB := strings.ReplaceAll(tB, "/", "-")

	FormatedTimeB, err := time.Parse(DBTimeFormat, timeB)

	if err != nil {
		return false
	}

	return FormatedTimeA.Equal(FormatedTimeB)
}
