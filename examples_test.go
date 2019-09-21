package tdigest

import "fmt"

func ExampleTD() {

	// Create a TDigest structure
	td := NewTD(nil)

	// Add data points ...
	for i := 0; i < 10; i++ {
		td.Add(float64(i))
	}

	td.Dump()
	// Output:
	// T-Digest (TD) structure description
	// Need digesting ? : false
	// Count  : 10
	// Mean   : 4.50
	// Min    : 0.00
	// Max    : 9.00
	// There are 10 buckets ...
	// 0 : Bucket center = 0.00, count = 1, count-below = 0, quartiles :0.050000
	// 1 : Bucket center = 1.00, count = 1, count-below = 1, quartiles :0.150000
	// 2 : Bucket center = 2.00, count = 1, count-below = 2, quartiles :0.250000
	// 3 : Bucket center = 3.00, count = 1, count-below = 3, quartiles :0.350000
	// 4 : Bucket center = 4.00, count = 1, count-below = 4, quartiles :0.450000
	// 5 : Bucket center = 5.00, count = 1, count-below = 5, quartiles :0.550000
	// 6 : Bucket center = 6.00, count = 1, count-below = 6, quartiles :0.650000
	// 7 : Bucket center = 7.00, count = 1, count-below = 7, quartiles :0.750000
	// 8 : Bucket center = 8.00, count = 1, count-below = 8, quartiles :0.850000
	// 9 : Bucket center = 9.00, count = 1, count-below = 9, quartiles :0.950000

}

func ExampleSizer() {

	td := NewTD(PolySizer(1.))

	for i := 0; i <= 10000; i++ {
		td.Add(float64(i))
	}
	fmt.Printf("\nMedian : %.3f", td.At(.5))
	fmt.Printf("\nQuartile : %.4f", td.Quartile(3000.))
	fmt.Printf("\nNb bkts  : %d", td.Size())

	// Output:
	// Median : 5000.000
	// Quartile : 0.3000
	// Nb bkts  : 21

}
