package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}

	return total/float64(len(xs))
}

// Find the smallest element in a slice.
func Min(xs []float64) float64 {
	
	min := xs[0]

	for _, val := range xs {
		if val < min {
			min = val
		}
	}

	return min
}

// Find the largest element in a slice.
func Max(xs []float64) float64 {

	max := xs[0]

	for _, val := range xs {
		if val > max {
			max = val
		}
	}

	return max
}
