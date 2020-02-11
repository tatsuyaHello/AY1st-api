package util

import "reflect"

// Contains slice に `fn` で定義した条件に一致する要素が含まれるかを返します
// `n` には `len(slice)` を指定します
func Contains(n int, fn func(i int) bool) bool {
	for i := 0; i < n; i++ {
		if fn(i) {
			return true
		}
	}
	return false
}

// FindIndex は slice 内の `fn` で定義した条件に一致する要素の index を返します
// `n` には `len(slice)` を指定します
func FindIndex(n int, fn func(i int) bool) (index int, found bool) {
	for i := 0; i < n; i++ {
		if fn(i) {
			return i, true
		}
	}
	return -1, false
}

// ListIDs は slice 内の `fn` で指定したIDのsliceを返します
// `n` には `len(slice)` を指定します
func ListIDs(n int, fn func(i int) uint64) []uint64 {
	list := make([]uint64, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, fn(i))
	}

	return list
}

// ListIDsOmitZero は slice 内の `fn` で指定した0以外のIDのsliceを返します
// `n` には `len(slice)` を指定します
func ListIDsOmitZero(n int, fn func(i int) uint64) []uint64 {
	list := make([]uint64, 0, n)
	for i := 0; i < n; i++ {
		if fn(i) == 0 {
			continue
		}
		list = append(list, fn(i))
	}

	return list
}

// ListStrings は slice 内の `fn` で指定したIDのsliceを返します
// `n` には `len(slice)` を指定します
func ListStrings(n int, fn func(i int) string) []string {
	list := make([]string, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, fn(i))
	}

	return list
}

// GetUniqueIDList 一意のID一覧を取得
func GetUniqueIDList(list []uint64) []uint64 {
	uniqueList := make([]uint64, 0, len(list))
	for _, id := range list {
		if Contains(len(uniqueList), func(i int) bool { return uniqueList[i] == id }) {
			continue
		}
		uniqueList = append(uniqueList, id)
	}
	return uniqueList
}

// GetUniqueIndexListStable 一意の一覧を順番を変更せずに取得
// 2重ループで処理するため、要素数が多くなると遅くなるが、ソート不可な対象に使用する
func GetUniqueIndexListStable(sliceList interface{}, equal func(i, j int) bool) []int {
	rv := reflect.ValueOf(sliceList)
	len := rv.Len()
	uniqueIndexes := make([]int, 0, len)
	for i := 0; i < len; i++ {
		found := false
		for j := 0; j < len; j++ {
			if i == j {
				continue
			}

			found = equal(i, j)
			if found {
				break
			}
		}
		if !found {
			uniqueIndexes = append(uniqueIndexes, i)
		}
	}
	return uniqueIndexes
}

// MaxIndex `gtFunc` で比較して最大の要素のindexを取得
func MaxIndex(n int, gtFunc func(a, b int) bool) int {
	if n < 1 {
		return -1
	}

	if n == 1 {
		return 0
	}

	j := 0

	for i := 1; i < n; i++ {
		if gtFunc(i, j) {
			j = i
		}
	}

	return j
}
