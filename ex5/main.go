package main

import "github.com/danteay/meetup-go-2020-10-15/ex5/print"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	print.Line(12)
	print.Line(12.34)
	print.Line("Jane")
	print.Line(ID("123-456-789"))
	print.Line(Person{name: "John Dou"})
	print.Line(make(chan bool))
	print.Line(&Person{name: "John Dou"})
}
