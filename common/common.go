package common

import "reflect"

type ViewField struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Hidden   bool   `json:"hidden"`
	Comment  string `json:"comment"`
	Format   string `json:"format,omitempty"`
	Analyzer string `json:"analyzer,omitempty"`

	Path []string `json:"-"`
}

func IsSlice(i any) bool {
	kind := reflect.ValueOf(i).Kind()
	return kind == reflect.Slice || kind == reflect.Array
}

func SliceLen(i any) int {
	val := reflect.Indirect(reflect.ValueOf(i))
	return val.Len()
}

func IsSameType(arr []any) bool {
	if len(arr) == 0 {
		return true
	}

	firstType := reflect.TypeOf(arr[0])
	for _, v := range arr {
		if reflect.TypeOf(v) != firstType {
			return false
		}
	}

	return true
}
