package print

import (
	"fmt"
	"reflect"
	"strings"
)

func Struct(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	switch t.Kind() {
	case reflect.Struct:
		s := structInfo(x)
		println(s)
	case reflect.Ptr:
		s := structInfo(reflect.Indirect(v).Interface())
		println("&" + s)
	default:
		fmt.Printf("err: unknown type '%v'\n", t)
	}
}

func structInfo(x interface{}) string {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	if t.Kind() != reflect.Struct {
		fmt.Printf("err: x %v value is not a struct\n", t)
		return ""
	}

	n := t.NumField() // panics if t.Kind != reflect.Struct

	sb := &strings.Builder{}

	fmt.Fprintf(sb, "%v{", t)

	for i := 0; i < n; i++ {
		tf := t.Field(i)
		vf := v.Field(i)

		fmt.Fprintf(sb, "%v: %v", tf.Name, vf)
	}

	fmt.Fprintf(sb, "}")

	return sb.String()
}
