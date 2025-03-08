package package1

import "fmt"

func init() {
	fmt.Println("Package 1 , init function 1")
}

func init() {
	fmt.Println("Package 1 , init function 2")
}

func P1() {
	fmt.Println("Package 1, P1 function")
}

func init() {
	fmt.Println("Package 1 , init function 3")
}
