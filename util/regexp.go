package util

import "strings"

// GetRegexpStringMatchForward 複数前方一致の文字列を取得
func GetRegexpStringMatchForward(element []string) string {
	regexpPatterns := make([]string, 0, len(element))
	for _, site := range element {
		regexpPatterns = append(regexpPatterns, "^"+site+".*")
	}
	return strings.Join(regexpPatterns, "|")
}

// GetRegexpStringMatchMid 複数中間一致の文字列を取得
func GetRegexpStringMatchMid(element []string) string {
	regexpPatterns := make([]string, 0, len(element))
	for _, site := range element {
		regexpPatterns = append(regexpPatterns, "^.*"+site+".*")
	}
	return strings.Join(regexpPatterns, "|")
}
