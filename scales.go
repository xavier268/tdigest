package tdigest

// Sizer defines the max size of a bucket for the provided quartiles limits.
type Sizer func(n int, ql, qr float64) float64
