package tdigest

import (
	"fmt"
	"testing"
)

func TestDigest3(t *testing.T) {

	// Define internal sizer
	sz := func(int, float64, float64) float64 {
		return 3 // max size per bucket
	}

	td := NewTD(sz)
	if td == nil || td.Count() != 0 || len(td.bkts) != 0 {
		t.Fatal(td)
	}
	td.Add(1., 7., 3., 2., 5., 0., 6.)
	//td.Dump()

	fmt.Println("Sorting ...")
	td.Sort()
	//td.Dump()

	fmt.Println("Digesting with sz3")
	td.digest()
	td.Dump()

	if td.Count() != 7 || len(td.bkts) != 3 {
		t.Fatal("Unexpected digest result with bcuket of fixed size 3")
	}
}

func TestDigest0(t *testing.T) {

	// Define internal sizer
	sz := func(int, float64, float64) float64 {
		return 0 // max size per bucket
	}

	td := NewTD(sz)
	td.Dump()
	if td == nil || td.Count() != 0 || len(td.bkts) != 0 {
		t.Fatal(td)
	}
	td.Add(1., 7., 3., 2., 5., 0., 6.)
	//td.Dump()

	fmt.Println("Sorting ...")
	td.Sort()
	//td.Dump()

	fmt.Println("Digesting with sz0")
	td.digest()
	//td.Dump()

	if td.Count() != 7 || len(td.bkts) != 7 {
		t.Fatal("Unexpected digest result with bcuket of fixed size 0")
	}
}
