package print

import (
	"fmt"
	"reflect"
)

func Line(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	switch t.Kind() {
	case reflect.Bool:
		fmt.Printf("boolean: %v\n", v)
	case reflect.Int:
		fmt.Printf("int: %v\n", v)
	case reflect.Float64:
		fmt.Printf("float64: %v\n", v)
	case reflect.String:
		fmt.Printf("string: %v\n", v)
	case reflect.Struct:
		structInfo(x)
	case reflect.Chan:
		fmt.Printf("chan: %v\n", v)
	case reflect.Ptr:
		fmt.Printf("pointer: %v\n", v)
	default:
		fmt.Printf("err: unknown type '%v'", t)
	}
}

func structInfo(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	if t.Kind() != reflect.Struct {
		fmt.Printf("err: x %v value is not a struct\n", t)
		return
	}

	n := t.NumField() // panics if t.Kind != reflect.Struct

	fmt.Printf("%v{", t)

	for i := 0; i < n; i++ {
		tf := t.Field(i)
		vf := v.Field(i)

		fmt.Printf("%v: %v", tf.Name, vf)
	}

	fmt.Println("}")
}
