package tdigest

// At provides the value corresponding to the provided quartile.
// The value returned is the best estimated such that the percentage
// of datapoints that are strictly below this value is the quartile.
// In other words, the ql matches the quartile.
// Panic if quartile is out of range 0.0-1.0
func (td *TD) At(quartile float64) (value float64) {

	if quartile < 0. || quartile > 1. {
		panic("Quartile should be between 0.0 and 1.0")
	}

	var q1, q2, v1, v2 float64
	// First, lets find the closest values
	for _, b := range td.bkts {
		if q2 = b.ql(td.n); q2 > quartile {
			// not yet ...
			q1 = q2
			v1 = b.mean()
		} else {
			// We crossed the value ...
			v2 = b.mean()
		}
	}

	if q1 == q2 {
		return v1
	}
	return v1 + (v2-v1)/(q2-q1)

}

// Quartile provides the quartile (percentage of data points that are
// equals or below that value).
func (td *TD) Quartile(value float64) float64 {
	return 0.
}

// Min provides an estimated minimum.
// Exact value only if first bucket contains a single element.
func (td *TD) Min() float64 {
	return td.bkts[0].mean()
}

// Max estimated maximum.
func (td *TD) Max() float64 {
	return td.bkts[len(td.bkts)-1].mean()
}

// Count gets the total number of data points seen
func (td *TD) Count() int {
	return td.n
}

// Mean is the (exact) mean of the datat set.
func (td *TD) Mean() float64 {
	s := 0.0
	for _, b := range td.bkts {
		s += b.sx
	}
	return s / float64(td.n)
}
