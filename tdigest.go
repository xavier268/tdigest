package tdigest

import (
	"fmt"
	"sort"
)

// TD is the TDigest type providing public access to statistical results.
type TD struct {
	n     int   // total number of data points
	bkts  []bkt // array of clusters
	sizer Sizer // the bucket sizer
}

// NewTD creates a new TDigest structure.
// If sizer is nil, all buckets will be sized 1.
func NewTD(sizer Sizer) *TD {
	td := new(TD)
	if sizer != nil {
		td.sizer = sizer
	} else {
		// This sizer prevents any digest.
		td.sizer = NilSizer
	}
	return td
}

// Digest will sort and digest, then return the object for chaining.
func (td *TD) Digest() *TD {
	td.sort().digest()
	return td
}

// digest is the core processing part of the algorithm.
// It examines all buckets, left to right, merging buckets every time it is possible.
// It assumes buckets are already ordered, and will leave them ordered.
func (td *TD) digest() *TD {

	for i := 0; i+1 < len(td.bkts); {
		// we consider merging (i) and (i+1)
		n := td.n
		q := td.bkts[i].q(n)
		sn := float64(td.bkts[i].n + td.bkts[i+1].n)
		if sn <= td.sizer(n, q) {
			// do the merge  and keep same i index ...
			td.bkts[i].merge(&td.bkts[i+1])
			// remove bucket i+1 - is this check necessary ?
			if i+2 < len(td.bkts) {
				td.bkts = append(td.bkts[:i+1], td.bkts[i+2:]...)
			} else {
				td.bkts = td.bkts[:i+1]
			}
			// do NOT increment i and loop to try to merge next bucket ...
		} else {
			// increment i ...
			i++
		}

	}
	return td

}

// Sort will sort the buckets.
// It will NOT merge them.
// It updates their sn field and the total count.
func (td *TD) sort() *TD {
	// sort ...
	sort.Slice(td.bkts, func(i, j int) bool {
		return td.bkts[i].mean() < td.bkts[j].mean()
	})
	// update the sn fields ...
	c := 0
	for i, b := range td.bkts {
		td.bkts[i].sn = c
		c += b.n
	}
	td.n = c
	return td
}

// Add a set of values.
// Values are added as independant buckets.
// You NEED to sort/digest when finished : this is not done by default.
func (td *TD) Add(values ...float64) *TD {
	for _, v := range values {
		b := bkt{sx: v, n: 1}
		td.bkts = append(td.bkts, b)
	}
	return td
}

// Merge will merge tt into td.
// It will combine the sizer, to the max of both.
// It will (Sort and) Digest automatically.
// tt is left unchanged.
func (td *TD) Merge(tt *TD) *TD {
	td.sizer = MaxSizer(td.sizer, tt.sizer)
	td.n += tt.n
	td.bkts = append(td.bkts, tt.bkts...)
	td.Digest()
	return td
}

// String display human readable value.
func (td *TD) String() string {
	s := fmt.Sprintf("\nT-Digest (TD) structure description"+
		"\nCount  : %d"+
		"\nMean   : %.2f"+
		"\nMin    : %.2f"+
		"\nMax    : %.2f"+
		"\nThere are %d buckets ..."+
		"\n",
		td.n,
		td.Mean(),
		td.Min(),
		td.Max(),
		len(td.bkts),
	)
	for i, b := range td.bkts {
		s += fmt.Sprintf("%d : %s, quartiles :%0.6f\n",
			i, b.String(), b.q(td.n))
	}

	return s
}

// Dump content in human readable format.
func (td *TD) Dump() {
	fmt.Println(td.String())
}
