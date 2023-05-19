package main

import "fmt"

func main() {
	// Create
	var stack []string
	// Push
	stack = append(stack, "world!")
	stack = append(stack, "Hello ")
	// Output: Hello world!
	fmt.Println(fmt.Sprint(stack[:1]) + fmt.Sprint(stack[1:]))

}
