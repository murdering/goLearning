package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := "hello world"
	fmt.Println(typeof(v))
	fmt.Println(typeOf(v))
}

func typeof(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	//... etc
	case string:
		return "string"
	default:
		_ = t
		return "unknown"
	}
}

func typeOf(v interface{}) string {
	return reflect.TypeOf(v).String()
}
