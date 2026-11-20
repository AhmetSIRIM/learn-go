package main

import "fmt"

func main() {
	var inferredTypeVariable = "Hello World!"
	fmt.Println(inferredTypeVariable)
	inferredTypeVariable = "I can be changed: every Go variable can, only a const cannot"
	fmt.Println(inferredTypeVariable)

	const constantValue = "I cannot be changed: I am a const"
	fmt.Println(constantValue)
	// constantValue = "new value" // Does not compile: cannot assign to constantValue

	fmt.Println("---")

	var explicitlyTypedVariable float64 = 1 // Typed on purpose: without float64, 1 would be an int
	var inferredIntVariable = 1             // No type given: 1 defaults to int
	fmt.Printf("%v is a %T\n", explicitlyTypedVariable, explicitlyTypedVariable)
	fmt.Printf("%v is an %T\n", inferredIntVariable, inferredIntVariable)
}
