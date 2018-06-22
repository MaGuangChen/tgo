package util

import "reflect"

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
