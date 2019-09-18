package tdigest

import "fmt"

// bkt aggregates a group of data points in a bucket.
type bkt struct {
	sx float64 // sum of values
	n  int     // number of data points in bucket
	sn int     // number of data points immediately in buckets before
	//            this bucket, excluding the points in the bucket
}

func (b *bkt) String() string {
	return fmt.Sprintf("Bucket center = %.2f, count = %d, count-below = %d",
		b.mean(), b.n, b.sn)
}

// add a value in the bucket, then return the updated bucket.
// Mostly used for debugging.
func (b *bkt) add(v float64) *bkt {
	b.sx += v
	b.n++
	// sn unchanged
	return b
}

// merge bb into b.
// bb is expected to have a mean that is greater or equal to b,
func (b *bkt) merge(bb *bkt) {
	if bb.mean() < b.mean() {
		panic("Bucket wrong ordering - cannot merge ! ")
	}
	if b.sn+b.n != bb.sn {
		panic("Inconsistent sn values - cannot merge !")
	}
	b.sx += bb.sx
	b.n += bb.n
	// note : b.sn does not change : always same number of data points to the left !
}

// mean gets the centroid (the mean of the bucket)
func (b *bkt) mean() float64 {
	return b.sx / float64(b.n)
}

// ql is the quartile for the left limit of the bucket
// n is the total nb of data points
func (b *bkt) ql(n int) float64 {
	return float64(b.sn) / float64(n)
}

// qr is the quartile for the right limit of the bucket
// n is the total nb of data points
func (b *bkt) qr(n int) float64 {
	return float64(b.sn+b.n) / float64(n)
}

// q is the average quartile for the bucket
// n is the total nb of data points
func (b *bkt) q(n int) float64 {
	return (float64(b.sn) + float64(b.n)/2) / float64(n)
}

func (b *bkt) Equal(bb *bkt) bool {
	if bb == nil {
		return false
	}
	return b.sx == bb.sx &&
		b.n == bb.n &&
		b.sn == bb.sn
}
