package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func celsiusToFahrenheitAndKelvin(c float64) (float64, float64) {
	return celsiusToFahrenheit(c), c + 273.15
}

func describeTemperature(celsius float64) string {
	switch { // Cases run top to bottom and the first true one wins, so each case only checks its upper bound
	case celsius <= 0:
		return "freezing"
	case celsius < 15:
		return "cold"
	case celsius <= 25:
		return "mild"
	case celsius < 100:
		return "hot"
	default: // The compiler does not check that the cases cover every value, so without default it reports "missing return"
		return "boiling"
	}
}

func showDeferOrder() {
	fmt.Println("start")
	defer fmt.Println("deferred first")  // Runs when showDeferOrder returns, not here
	defer fmt.Println("deferred second") // Deferred calls stack up: the last one deferred runs first
	fmt.Println("end")
}

func showDeferArgumentTiming() {
	count := 1
	fmt.Println("count before defer:", count)
	defer fmt.Println("deferred print, runs last, shows the count saved at defer time:", count) // count is read on this line, not when the print runs
	count = 2
	fmt.Println("count after change:", count)
}

func showNestedDeferStacks() {
	defer fmt.Println("outer: deferred 1")
	defer fmt.Println("outer: deferred 2") // A function can defer any number of calls: there is no fixed limit, only memory

	defer func() { // An anonymous function can be deferred too, for cleanup longer than one call; the () at the end calls it
		fmt.Println("outer: deferred 3, line 1")
		fmt.Println("outer: deferred 3, line 2")
	}()

	fmt.Println("outer: calls inner")

	func() { // An anonymous function has its own defer stack, which empties when it returns, before the outer one
		defer fmt.Println("  inner: deferred A")
		defer fmt.Println("  inner: deferred B")
		fmt.Println("  inner: body")
	}()

	fmt.Println("outer: back from inner")
}
