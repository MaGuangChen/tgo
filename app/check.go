package app

import (
	"fmt"
	"reflect"
)

// CheckType : 檢查型別
func CheckType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// CheckError : 檢查錯誤
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error is not nil plz check: ", err)
	}
}
