package simpleinterest

func CalculateSimpleInterest(p float64, r float64, t float64) float64 {
	var interest float64 = (p * r * t) / 100

	return interest
}
