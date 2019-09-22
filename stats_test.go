package tdigest

import (
	"math"
	"testing"
)

// Single value buckets
func SetUp0() *TD {
	td := NewTD(nil)
	data := make([]float64, 100)
	for i := 0; i < 100; i++ {
		data[i] = float64(i)
	}
	td.Add(data...)
	return td
}

// Setup for realistic sizer
func SetUp1() *TD {
	td := NewTD(PolySizer(1.0))
	v := make([]float64, 10000000)
	for i := 0; i < 10000000; i++ {
		v[i] = float64((i * 98013) % 1000000)
	}
	td.Add(v...)
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
