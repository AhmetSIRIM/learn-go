package main

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func celsiusToFahrenheitAndKelvin(c float64) (float64, float64) {
	return celsiusToFahrenheit(c), c + 273.15
}
