package main

import "fmt"

func sayHello() string {
	return "Hello world!"
}

func main() {
	hello := sayHello()
	fmt.Println(hello)
}
