package util

import "reflect"

// GetStructTagValueList Struct Field から、指定の tagName の一覧を取得
func GetStructTagValueList(i interface{}, tagName string) map[string]string {

	st := reflect.TypeOf(i)

	if i == nil {
		return map[string]string{}
	}

	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}
	if st.Kind() != reflect.Struct {
		panic("input object is not struct")
	}

	tagValueList := make(map[string]string, st.NumField())

	for i := 0; i < st.NumField(); i++ {
		sf := st.Field(i)
		v := sf.Tag.Get(tagName)
		if v != "" {
			tagValueList[sf.Name] = v
		}
	}

	return tagValueList
}
