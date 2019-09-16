package tdigest

// At provides the value corresponding to the provided quartile.
// Panic if quartile is out of range 0.0-1.0
func (td *TD) At(quartile float64) (value float64) {
	// TO DO - do we need min/max in TD ?
	// Probably not if we force extreme quartile to be sized of 1 to capture min/max ?
	return 0.
}

// Quartile provides the quartile (percentage of data points that are
// equals or below that value).
func (td *TD) Quartile(value float64) float64 {
	// TO DO
	return 0.
}
