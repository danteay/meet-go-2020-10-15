package main

import "github.com/danteay/meetup-golang-2020-10-15/ex6/print"

type Person struct {
	name string
}

func main() {
	print.Struct(Person{name: "John Dou"})
	print.Struct(&Person{name: "John Dou"})
}
