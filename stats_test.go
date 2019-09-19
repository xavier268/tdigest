package tdigest

import (
	"math"
	"testing"
)

// Single value buckets
func SetUp0() *TD {
	td := NewTD(nil)
	for i := 0; i < 100; i++ {
		td.Add(float64(i))
	}
	td.Digest()
	return td
}

// Setup for realistic sizer
func SetUp1() *TD {
	td := NewTD(PolySizer(1.0))
	for i := 0; i < 10000000; i++ {
		td.Add(float64((i * 98013) % 1000000))
		if i%1000 == 0 { // Limit memory footprint
			td.Digest()
		}
	}
	td.Digest()
	return td
}

func TestBasicStats(t *testing.T) {
	td := SetUp0()
	// td.Dump()
	if td.Count() != 100 || td.Min() != 0.0 || td.Max() != 99.0 || td.Mean() != 49.5 {
		td.Dump()
		t.FailNow()
	}
}

func TestStat(t *testing.T) {
	// t.Skip()
	td := SetUp1()
	td.Dump()
	if td.Count() != 10000000 ||
		math.Abs(td.Min()) > 0.00001 ||
		math.Abs(td.Max()-999999.0) > 0.00001 ||
		math.Abs(td.Mean()-499999.5) > 0.00001 {
		td.Dump()
		t.FailNow()
	}

}
