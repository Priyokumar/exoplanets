package gravity

type GravityCalculator interface {
	Calculate() (float64, error)
}

func CalculateGravity(c GravityCalculator) (float64, error) {
	return c.Calculate()
}
