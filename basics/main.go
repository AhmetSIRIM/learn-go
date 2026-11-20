package main

import (
	"fmt"
	"log"
	"math"
	"math/rand/v2"
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

	fmt.Println("\n********************************************************************************************")
	fmt.Println()

	var celsius float64 = 100
	fmt.Println(celsius)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Println("fahrenheit: " + strconv.FormatFloat(fahrenheit, 'f', -1, 64))
	fahrenheit, kelvin := celsiusToFahrenheitAndKelvin(celsius)
	fmt.Println("fahrenheit and kelvin:", fahrenheit, "&", kelvin)

	fmt.Println("\n********************************************************************************************")
	fmt.Println()

	start := -43
	end := -37

	for index := start; index <= end; index++ { // Three-part form: init; condition; post
		fmt.Printf("c: %d & f: %.1f\n", index, celsiusToFahrenheit(float64(index))) // %.1f: one decimal, hides float rounding noise such as -45.400000000000006
	}

	fmt.Print("\n---\n\n")

	for index := range 10 { // Range over an int: index runs from 0 to 9, the same as index := 0; index < 10; index++
		fahrenheit, kelvin = celsiusToFahrenheitAndKelvin(float64(index)) // Two results cannot sit next to other arguments, so they are taken apart first
		fmt.Printf("c: %d & f: %.1f & k: %.2f\n", index, fahrenheit, kelvin)
	}

	fmt.Print("\n---\n\n")

	word := "go"
	for len(word) < 20 { // While form: only a condition, for a loop whose step count is not known up front
		fmt.Println(len(word), word)
		word += word // Doubles the word, so its length goes 2, 4, 8, 16, 32
	}
	fmt.Println(len(word), word)

	fmt.Print("\n********************************************************************************************\n\n")

	number := rand.Float64() * 9 // A positive number below 9
	if rand.IntN(2) == 0 {       // A coin flip: IntN(2) returns 0 or 1, and math/rand has no Bool
		number = -number
	}

	if newtonRaphsonRoot, err := newtonRaphsonSquareRoot(number, true); err != nil { // Short statement: newtonRaphsonRoot and err live only in this if and its else
		log.Println(err) // log writes to stderr, the channel for errors, so the failure is never silent
	} else {
		fmt.Print("\n")
		fmt.Printf("'math.Sqrt'            :%.1f\nnewtonRaphsonSquareRoot:%.1f\n", math.Sqrt(number), newtonRaphsonRoot)
	}

	fmt.Print("\n********************************************************************************************\n\n")

	fmt.Println("-5:", describeTemperature(-5))
	fmt.Println("0:", describeTemperature(0)) // Boundary
	fmt.Println("10:", describeTemperature(10))
	fmt.Println("15:", describeTemperature(15)) // Boundary
	fmt.Println("25:", describeTemperature(25)) // Boundary
	fmt.Println("60:", describeTemperature(60))
	fmt.Println("100:", describeTemperature(100)) // Boundary

	fmt.Print("\n********************************************************************************************\n\n")

	showDeferOrder()
	fmt.Println("---")
	showDeferArgumentTiming()
	fmt.Println("---")
	showNestedDeferStacks()
}
