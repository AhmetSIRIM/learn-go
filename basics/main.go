package main

import (
	"fmt"
	"strconv"
)

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

	fmt.Println("---")

	var zeroValueOfFloat64 float64
	convertedFloat64 := strconv.FormatFloat(zeroValueOfFloat64, 'f', -1, 64) // 'f': no exponent, -1: fewest digits that parse back to the same value, 64: the value is a float64
	fmt.Println("Zero value of float64: " + convertedFloat64)                // + joins only strings: Go never converts a number to a string on its own
	var zeroValueOfString string
	fmt.Println("Zero value of string: " + zeroValueOfString)
	fmt.Printf("Zero value of string, quoted: %q\n", zeroValueOfString) // %q puts quotes around it, so the empty string shows as ""
	var zeroValueOfBool bool
	fmt.Println("Zero value of bool: " + strconv.FormatBool(zeroValueOfBool))
	var zeroValueOfInt int
	fmt.Println("Zero value of int:", zeroValueOfInt) // Println takes values of any type, so no strconv is needed
}
