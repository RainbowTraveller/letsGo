package main

import "fmt"

type Dog struct {
	name string
}

type Cat struct {
	age int
}

func main() {
	var empty interface{}
	empty = Dog{}
	empty.name = "Marshal"
	fmt.Println(empty)
	empty = Cat{}
	empty.age = 4
	fmt.Println(empty)
}
