package tdigest

// Sizer defines the max size of a bucket for the provided quartiles average.
// Assumptions are that 0 <= q <= 1.0
type Sizer func(n int, q float64) float64

// =========  High Order Functions ==================
// These functions are used to generate custom Sizer.
// ==================================================

// NilSizer : When sizer is not set, use this value.
func NilSizer() Sizer {
	return MakeConstSizer(1)
}

// PolySizer defines a polynom structure for the sizer
// Use scale to adjust the numer of buckets.
// Number of bucket increase when scale increase.
func PolySizer(scale float64) Sizer {
	return func(n int, q float64) float64 {
		return float64(n) * q * (1 - q) / scale
	}
}

// LinearSizer will size buckets linearly, ending up with appx 2 to 3  buckets.
// Use scale to adjust actual number of buckets.
// Number of bucket increase when scale increase.
func LinearSizer(scale float64) Sizer {
	return func(n int, _ float64) float64 {
		return float64(n) / scale
	}
}

// MakeConstSizer provides a Sizer of a fixed value, k.
func MakeConstSizer(k int) Sizer {
	return func(_ int, _ float64) float64 {
		return float64(k)
	}
}

// ScaleSizer applies a scale to an existing sizer
// Typically, scale is proportional to the number of buckets.
func ScaleSizer(s Sizer, scale float64) Sizer {
	return func(n int, q float64) float64 {
		return s(n, q) / scale
	}
}

// MaxSizer returns a new Sizer that provides the max between the s1 and s2 Sizers.
// Used when merging 2 tdigests.
func MaxSizer(s1, s2 Sizer) Sizer {
	return func(n int, q float64) float64 {
		v1 := s1(n, q)
		v2 := s2(n, q)
		if v1 <= v2 {
			return v2
		}
		return v1
	}
}
