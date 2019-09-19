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

	var v1, q1 float64

	// First, lets find the closest bucket immediately below
	for i, b := range td.bkts {
		if b.q(td.n) < quartile {
			// not yet, keep moving but get values ...
			v1 = b.mean()
			q1 = b.q(td.n)
		} else {
			// We found a  bucket !
			if i == 0 { // no previous, don't interpolate !
				return b.mean()
			}
			// Interpolate with previous
			v := b.mean()
			q := b.q(td.n)
			return v1 + (quartile-q1)*(v-v1)/(q-q1)

		}
	}

	panic("No value matching requested quartile - unexpected error !")

}

// Quartile provides the quartile (percentage of data points that are
// equals or below that value).
func (td *TD) Quartile(value float64) float64 {

	var v1, q1 float64

	// First, lets find the closest bucket immediately below
	for i, b := range td.bkts {
		if b.mean() < value {
			// not yet, keep moving but get values ...
			v1 = b.mean()
			q1 = b.q(td.n)
		} else {
			// We found a  bucket !
			if i == 0 { // no previous, don't interpolate !
				return 0.0
			}
			// Interpolate with previous
			v := b.mean()
			q := b.q(td.n)
			return q1 + (value-v1)*(q-q1)/(v-v1)
		}
	}
	// beyond the last bucket ...
	return 1.0
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
