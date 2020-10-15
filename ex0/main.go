package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	fmt.Println(12)
	fmt.Println(12.34)
	fmt.Println("John")
	fmt.Println(true)
	fmt.Println(Person{})
	fmt.Println(&Person{})
	fmt.Println([]string{})
	fmt.Println(make(chan bool))
}
