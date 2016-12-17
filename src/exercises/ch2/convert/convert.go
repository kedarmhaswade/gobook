package convert

func MeterToFeet(x float64) float64 {
	return x * 3.2795
}

func FeetToMeter(x float64) float64 {
	return x * 1 / MeterToFeet(1)
}

func FeetToInch (x float64) float64 {
	return 12 * x
}

func InchToFeet(x float64) float64 {
	return x * 1 / FeetToInch(1)
}

func MeterToInch(x float64) float64 {
	return MeterToFeet(x) * FeetToInch(x)
}

func InchToMeter(x float64) float64 {
	return x * 1/MeterToInch(1)
}

func PoundsToGram(x float64) float64 {
	return x * 454
}

func GramToPounds(x float64) float64 {
	return x * 1 / PoundsToGram(1)
}
func CelsiusToFahrenheit(x float64) float64 {
	return (x * 9) / 5 + 32
	// is there a good way to share a function like FahrenheitToCelsius (not easily invertible)?
}
func FahrenheitToCelsius(x float64) float64 {
	return (x - 32) * 5 / 9
}
