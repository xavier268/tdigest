package tdigest

import (
	"testing"
)

func TestDigest0(t *testing.T) {

	// Define internal sizer
	sz := MakeConstSizer(0)

	td := NewTD(sz)
	if td == nil || td.Count() != 0 || len(td.bkts) != 0 {
		td.Dump()
		t.Fatal(td)
	}
	td.Add(1., 7., 3., 2., 5., 0., 6.)
	td.Digest()
	if td.Count() != 7 || len(td.bkts) != 7 {
		td.Dump()
		t.Fatal("Unexpected digest result with bcuket of fixed size 0")
	}
}

func TestDigest3(t *testing.T) {

	// Define internal sizer
	sz := MakeConstSizer(3)

	td := NewTD(sz)
	if td == nil || td.Count() != 0 || len(td.bkts) != 0 {
		t.Fatal(td)
	}
	td.Add(1., 7., 3., 2., 5., 0., 6.)
	td.Digest()

	if td.Count() != 7 || len(td.bkts) != 3 {
		td.Dump()
		t.Fatal("Unexpected digest result with bcuket of fixed size 3")
	}
}

func TestDigestLinear(t *testing.T) {
	td := NewTD(LinearSizer)
	td.Add(1., 7., 3., 2., 5., 0., 6.)
	td.Digest()
	//td.Dump()
	if td.Count() != 7 || len(td.bkts) != 3 {
		td.Dump()
		t.Fatal("Unexpected digest result with bcuket of Linear Size")
	}
}
