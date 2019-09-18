package tdigest

// Sizer defines the max size of a bucket for the provided quartiles limits.
// Assumptions are that 0 <= ql <= qr <= 1.0
type Sizer func(n int, ql, qr float64) float64

// compiler checks
var _ Sizer = LinearSizer
var _ Sizer = PolySizer

// NilSizer : When sizer is not set, use this value.
var NilSizer Sizer = MakeConstSizer(1)

// PolySizer defines a polynom structure for the sizer
// Use ScaleSizer to scale this to the target ,nbr of buckets.
func PolySizer(n int, ql, qr float64) float64 {
	return float64(n) * float64(n) * qr * ql * (1 - qr) * (1 - ql)
}

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

// ScaleSizer applies a scale to an existing sizer
// Typically, scale is same order of magnitue as the number of buckets.
func ScaleSizer(s Sizer, scale float64) Sizer {
	return func(n int, ql, qr float64) float64 {
		return s(n, ql, qr) / scale
	}
}

// MaxSizer returns a new Sizer that provides the max between the s1 and s2 Sizers.
// Used when merging 2 tdigests.
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

// ForceMinMax modify s to ensure the min/max buckets are sized 1.
// Other bucket sized are derived from s.
func ForceMinMax(s Sizer) Sizer {
	return func(n int, ql, qr float64) float64 {
		if ql <= 0.0 || qr >= 1. {
			return 1
		}
		return s(n, ql, qr)
	}
}

// ForceMinimumResolution modify s to ensure buckets never exceed a with of qres.
// Other bucket sized are derived from s.
func ForceMinimumResolution(s Sizer, qres float64) Sizer {
	return func(n int, ql, qr float64) float64 {
		if qr-ql <= 0. || qr-ql >= qres {
			return 1
		}
		return s(n, ql, qr)
	}
}

// FromScaleFunction drives a Sizer from a scale function as defined in the T.Dunning paper.
func FromScaleFunction(scale func(int, float64) float64) Sizer {
	return func(n int, ql, qr float64) float64 {
		return scale(n, qr) - scale(n, ql)
	}
}
