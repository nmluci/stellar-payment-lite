package structutil

import (
	"reflect"
	"strconv"
)

func ParseTag[T any]() []string {
	var t T
	res := []string{}

	baseType := reflect.TypeOf(t)
	for i := 0; i < baseType.NumField(); i++ {
		if col := baseType.Field(i).Tag.Get("db"); col != "" && col != "-" {
			res = append(res, col)
		}
	}

	return res
}

func StringToInt64(s string) (res int64) {
	res, _ = strconv.ParseInt(s, 10, 64)
	return
}

func StringToFloat32(s string) (res float32) {
	x, _ := strconv.ParseFloat(s, 32)
	return float32(x)
}
