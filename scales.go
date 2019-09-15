package tdigest

// Sizer defines the max size of a bucket for the provided quartiles limits.
// Assumptions are that 0 <= ql <= qr <= 1.0
type Sizer func(n int, ql, qr float64) float64

// compiler checks
var _ Sizer = LinearSizer

// LinearSizer will size buckets linearly, ending up with appx 2 to 3  buckets.
func LinearSizer(n int, ql, qr float64) float64 {
	return float64(n) / 2.
}
