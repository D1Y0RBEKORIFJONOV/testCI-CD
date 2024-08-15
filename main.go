package main

import "fmt"

func main() {
	msg := sayHelloWorld("Test")
	fmt.Println(msg)
}

func sayHelloWorld(msg string) string {
	return "Hello, " + msg
}
