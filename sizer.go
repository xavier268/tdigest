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

// =========  High Order Functions ==================

// MakeConstSizer provides a Sizer of a fixed value, k.
func MakeConstSizer(k int) Sizer {
	return func(_ int, _, _ float64) float64 {
		return float64(k)
	}
}

// MaxSizer returns a new Sizer that provides the max between the s1 and s2 Sizers.
func MaxSizer(s1, s2 Sizer) Sizer {
	return func(n int, ql, qr float64) float64 {
		v1 := s1(n, ql, qr)
		v2 := s2(n, ql, qr)
		if v1 <= v2 {
			return v2
		}
		return v1
	}
}
