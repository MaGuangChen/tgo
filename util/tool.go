package util

import (
	"fmt"
	"reflect"
)

/* 檢查型別 */
func CheckType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

/* 檢查錯誤 */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error is not nil plz check: ", err)
	}
}
