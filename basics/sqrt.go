package main

import (
	"fmt"
	"math"
)

func newtonRaphsonSquareRoot(radicand float64, printSteps bool) (float64, error) {
	if radicand < 0 { // Guard clause: reject invalid input early, so the main path below stays unindented
		return 0, fmt.Errorf("cannot take the square root of a negative number: %v", radicand)
	}

	currentGuess := 1.0
	const tolerance = 1e-9

	for {
		previousGuess := currentGuess

		if printSteps {
			fmt.Printf("The guess is: %.2f\n", previousGuess)
		}

		currentGuess -= (currentGuess*currentGuess - radicand) / (2 * currentGuess)

		guessDifference := math.Abs(currentGuess - previousGuess)
		if guessDifference < tolerance {
			break
		}
	}

	return currentGuess, nil
}
