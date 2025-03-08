package main

import (
	"example/package1"
	"example/package2"
	"fmt"
)

// GLOBAL VARIABLES

var a string = "Hello"
var b string = "World"

func main() {
	package2.P2()
	fmt.Println("Hello")
	package1.P1()
	example()
	fmt.Println(a)
	fmt.Println(b)
}

func example() {
	fmt.Println("example")
}

func init() {
	fmt.Println("Package main, init function")
	a = "I love"
	b = "Golang"
}
