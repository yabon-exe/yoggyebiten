package physics

func MoveFall(v0 float64, g float64, t int) float64 {
	return v0 + (-g)*float64(t)
}
