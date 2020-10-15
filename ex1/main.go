package main

import "fmt"

func main() {
	var x interface{}

	x = 3.45
	fmt.Printf("x: type = %T, value = %v\n", x, x)

	y := x
	fmt.Printf("y: type = %T, value = %v\n", y, y)

	x = &struct{ name string }{}
	fmt.Printf("x: type = %T, value = %v\n", x, x)

	z := x
	fmt.Printf("z: type = %T, value = %v\n", z, z)
}
