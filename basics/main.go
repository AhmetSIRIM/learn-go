package main

import "fmt"

var iAmGlobalVariable = "I am a global variable" // Package level: var is required here, := only works inside a function

// shortAtPackageLevel := "x" // Does not compile: non-declaration statement outside function body

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

	fmt.Println("---")

	shortDeclaredVariable := "I was declared with :=" // Short declaration: var plus type inference, only inside a function

	fmt.Println(iAmGlobalVariable)
	fmt.Println(shortDeclaredVariable)
	shortDeclaredVariable = "I am the modified shortDeclaredVariable: " + shortDeclaredVariable // Assignment uses =, not :=
	fmt.Println(shortDeclaredVariable)
	// shortDeclaredVariable := "again" // Does not compile: no new variables on left side of :=
}
