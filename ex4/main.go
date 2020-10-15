package main

import (
	"fmt"
	"reflect"
)

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	structInfo(12.34)
	structInfo(Person{name: "John Doe"})
	structInfo(ID("John"))
}

func structInfo(x interface{}) {
	t := reflect.TypeOf(x)

	if t.Kind() != reflect.Struct {
		fmt.Printf("err: x %v value is not a struct\n", t)
		return
	}

	n := t.NumField() // panics if t.Kind != reflect.Struct

	for i := 0; i < n; i++ {
		f := t.Field(i)

		fmt.Printf("field: index = %v name = %v type = %v\n", i, f.Name, f.Type)
	}
}
