package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

// `math/rand` パッケージの `lockedSource` を引用
// https://golang.org/src/math/rand/rand.go?s=12478:12833
type lockedSource struct {
	lk  sync.Mutex
	src rand.Source64
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Uint64() (n uint64) {
	r.lk.Lock()
	n = r.src.Uint64()
	r.lk.Unlock()
	return
}

// DatetimeTo8CharactersPtr は '/' で区切られたDatetimeを8文字の文字列に変換
func DatetimeTo8CharactersPtr(ps *string) string {
	if ps == nil {
		return ""
	}

	s := *ps

	s = strings.ReplaceAll(s, "/", "")
	s = strings.ReplaceAll(s, "-", "")
	return s
}

// BoolTo01 はboolを"0"や"1"に変換して返す
func BoolTo01(b *bool) string {
	if b == nil {
		return "0"
	}

	if *b {
		return "1"
	}
	return "0"
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

var rng = rand.New(&lockedSource{
	src: rand.NewSource(time.Now().UnixNano()).(rand.Source64),
})

// RandPasswordString is created new password.
func RandPasswordString(n int) string {
	rs6Letters := "abcdefghjkmnpqrtuvwxy346789"
	rs6LetterIdxBits := 6
	rs6LetterIdxMask := (int64)(1<<rs6LetterIdxBits - 1)
	rs6LetterIdxMax := 63 / rs6LetterIdxBits

	b := make([]byte, n)
	cache, remain := rng.Int63(), rs6LetterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = rng.Int63(), rs6LetterIdxMax
		}
		idx := int(cache & rs6LetterIdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i--
		}
		cache >>= rs6LetterIdxBits
		remain--
	}
	password := string(b)
	if IsValidPassword(password) {
		return password
	}
	return RandPasswordString(n)
}

// IsValidPassword は正しいパスワードかチェック
// アルファベット(⼤⽂字⼩⽂字の制約なし)と半⾓数字の混合で、記号は含まない
func IsValidPassword(password string) bool {
	if IsNumericOnly(password) {
		return false
	}

	if IsAlphabeticOnly(password) {
		return false
	}

	return IsAlphaNumericOnly(password)
}

// IsNumericOnly は文字列が数字のみかチェック
func IsNumericOnly(s string) bool {
	for _, r := range s {
		if r < '0' || '9' < r {
			return false
		}
	}
	return true
}

// IsAlphabeticOnly は文字列がアルファベット(大文字,小文字)のみかチェック
func IsAlphabeticOnly(s string) bool {
	for _, r := range s {
		if (r < 'a' || 'z' < r) && (r < 'A' || 'Z' < r) {
			return false
		}
	}
	return true
}

// IsAlphaNumericOnly は文字列がアルファベット(大文字,小文字)、数字のみかチェック
func IsAlphaNumericOnly(s string) bool {
	for _, r := range s {
		if (r < 'a' || 'z' < r) && (r < 'A' || 'Z' < r) && (r < '0' || '9' < r) {
			return false
		}
	}
	return true
}

// PadLeftZero は0で左埋め
func PadLeftZero(len int, s string) string {
	format := "%0" + strconv.Itoa(len) + "s"
	return fmt.Sprintf(format, s)
}
