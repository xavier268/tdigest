package examples

import (
	"testing"

	"github.com/xavier268/tdigest"
)

// Naive additions, (non-efficient), for reference.
func BenchmarkAddOneByOne(b *testing.B) {

	td := tdigest.NewTD(tdigest.PolySizer(3.))
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {

		for j := 0; j < 10000; j++ {
			td.Add(float64(i))
		}

	}

}

// The following is about 50 x faster that single adds,
// for the exact same data added.
func BenchmarkAdd10000By10000(b *testing.B) {

	td := tdigest.NewTD(tdigest.PolySizer(3.))
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data := make([]float64, 10000)
		for j := 0; j < 10000; j++ {
			data[j] = float64(i)
		}
		td.Add(data...)
	}
}
