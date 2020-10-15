package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}

	x = 3.45

	t := reflect.TypeOf(x) // x.(type)
	v := reflect.ValueOf(x)

	fmt.Printf("x: type = %v, value = %v\n", t, v)

	y := x
	fmt.Printf("y: type = %T, value = %v\n", y, y)

	x = &struct{ name string }{}

	t = reflect.TypeOf(x)
	v = reflect.ValueOf(x)

	fmt.Printf("x: type = %v, value = %v\n", t, v)

	z := x
	fmt.Printf("z: type = %T, value = %v\n", z, z)
}
