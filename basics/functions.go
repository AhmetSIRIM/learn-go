package main

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
