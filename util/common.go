package util

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

/*
	判断是否为空
*/
func IsEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type())
}

/*
	数组去重
 */
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

/*
	大端法
 */
func ByteToInt(b []byte) (int, error) {
	bin_buf := bytes.NewBuffer(b)
	var x int
	err := binary.Read(bin_buf, binary.BigEndian, &x)
	if err != nil {
		return 0, err
	}
	return x, nil
}
