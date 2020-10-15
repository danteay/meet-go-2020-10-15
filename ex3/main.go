package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}

	x = &struct{ name string }{}

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("x: type = %v, value = %v\n", t, v)

	x = new(string)

	t1 := reflect.TypeOf(x)
	v1 := reflect.ValueOf(x)

	fmt.Printf("x: type = %v, value = %v\n", t1, v1)

	// obtain kind of type (kind is a supertype or category of a type)
	// kind tell us how we can treat the value of a variable
	fmt.Printf("t: type = %v, kind = %v\n", t, t.Kind())
	fmt.Printf("t1: type = %v, kind = %v\n", t1, t1.Kind())
}
