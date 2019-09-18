package tdigest

import "fmt"

func ExampleTD() {

	// Create a TDigest structure
	td := NewTD(nil)

	// Add data points ...
	for i := 0; i < 10; i++ {
		td.Add(float64(i))
	}

	// Final digest
	td.Digest()
	td.Dump()
	// Output:
	// T-Digest (TD) structure description
	// Count  : 10
	// Mean   : 4.50
	// Min    : 0.00
	// Max    : 9.00
	// There are 10 buckets ...
	// 0 : Bucket center = 0.00, count = 1, count-below = 0, quartiles left :0.000000 right:0.100000
	// 1 : Bucket center = 1.00, count = 1, count-below = 1, quartiles left :0.100000 right:0.200000
	// 2 : Bucket center = 2.00, count = 1, count-below = 2, quartiles left :0.200000 right:0.300000
	// 3 : Bucket center = 3.00, count = 1, count-below = 3, quartiles left :0.300000 right:0.400000
	// 4 : Bucket center = 4.00, count = 1, count-below = 4, quartiles left :0.400000 right:0.500000
	// 5 : Bucket center = 5.00, count = 1, count-below = 5, quartiles left :0.500000 right:0.600000
	// 6 : Bucket center = 6.00, count = 1, count-below = 6, quartiles left :0.600000 right:0.700000
	// 7 : Bucket center = 7.00, count = 1, count-below = 7, quartiles left :0.700000 right:0.800000
	// 8 : Bucket center = 8.00, count = 1, count-below = 8, quartiles left :0.800000 right:0.900000
	// 9 : Bucket center = 9.00, count = 1, count-below = 9, quartiles left :0.900000 right:1.000000
}

func ExampleSizer() {
	// A polynomial sizer, scaled to less than 10 buckets.
	sz := ScaleSizer(PolySizer, 60.)
	// Where min and max are protected by single element buckets
	sz = ForceMinMax(sz)

	td := NewTD(sz)

	for i := 0; i <= 10000; i++ {
		td.Add(float64(i))
		// Bounded memory footprint ...
		if i%100 == 0 {
			td.Digest()
		}
	}
	td.Digest()
	fmt.Printf("\nNumber of buckets : %d\nMedian : %.0f\n",
		len(td.bkts), td.At(.5))

	// Output:
	// interpolating around 20
	// Number of buckets : 38
	// Median : 5000
}
