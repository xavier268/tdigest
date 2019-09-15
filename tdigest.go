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
func NewTD(sizer Sizer) *TD {
	td := new(TD)
	td.sizer = sizer
	return td
}

// Count gets the total number of data points seen
func (td *TD) Count() int {
	return td.n
}

// digest is the core processing part of the algorithm.
// It examines all buckets, left to right, merging buckets every time it is possible.
// It assumes buckets are already ordered, and will leave them ordered.
func (td *TD) digest() {

	for i := 0; i+1 < len(td.bkts); {
		// we consider merging (i) and (i+1)
		n := td.n
		ql := td.bkts[i].ql(n)
		qr := td.bkts[i+1].qr(n)
		sn := float64(td.bkts[i].n + td.bkts[i+1].n)
		if sn <= td.sizer(n, qr, ql) {
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

}

// Sort will sort the buckets.
// It will NOT merge them.
// It updates their sn field and the total count.
func (td *TD) Sort() *TD {
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
// You need to sort/digest when finished : this is not done by default.
func (td *TD) Add(values ...float64) {
	for _, v := range values {
		b := bkt{sx: v, n: 1}
		td.bkts = append(td.bkts, b)
	}
}

// String display human readable value.
func (td *TD) String() string {
	s := fmt.Sprintf("\nT-Digest (TD) structure description\nTotal count   : %d\nTotal buckets : %d\nBuckets ... \n",
		td.n,
		len(td.bkts),
	)
	for _, b := range td.bkts {
		s += fmt.Sprintln(b.String())
	}
	return s
}

// Dump content in human readable format.
func (td *TD) Dump() {
	fmt.Println(td.String())
}
