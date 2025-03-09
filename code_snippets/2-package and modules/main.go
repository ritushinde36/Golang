package main

import (
	"application/utility"
	"fmt"
)

func main() {
	fmt.Println("Hello from the main package, main function")
	utility.Utility_func1()
	sayHello()
	utility.Utility_func2()
	fmt.Printf("My favourite language is %v", utility.Fav_language)

}

func sayHello() {
	fmt.Println("Hello from the main package, sayHello function")
}
