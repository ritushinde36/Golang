package package2

import "fmt"

func init() {
	fmt.Println("Package 2 , init function 1")
}

func init() {
	fmt.Println("Package 2 , init function 2")
}

func P2() {
	fmt.Println("Package 2, P2 function")
}
