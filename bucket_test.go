package tdigest

import (
	"testing"
)

func TestEqual(t *testing.T) {
	var b1, b2, b3 *bkt

	b1 = new(bkt).add(1.).add(2.)
	b2 = new(bkt)
	b3 = new(bkt).add(2.).add(1.)
	if b1.Equal(b2) {
		t.Fatal(b1, b2)
	}
	if !b1.Equal(b3) {
		t.Fatal(b1, b3)
	}
	if !b1.Equal(b1) {
		t.Fatal(b1)
	}
	if b1.Equal(nil) {
		t.Fatal(b1)
	}
}

func TestMerge0(t *testing.T) {
	var b1, b2, b3 *bkt

	b1 = new(bkt).add(1.).add(2.)
	b2 = new(bkt)
	b2.sn = 2

	b3 = new(bkt).add(2.).add(1.)
	b3.sn = 0

	b1.merge(b2)
	if !b1.Equal(b3) {
		t.Fatal(b1, b3)
	}
}
func TestMerge1(t *testing.T) {
	var b1, b2, b3 *bkt

	b1 = new(bkt).add(1.).add(2.)
	b1.sn = 10
	b2 = new(bkt).add(3.)
	b2.sn = 12

	b3 = new(bkt).add(2.).add(1.).add(3.)
	b3.sn = 10
	b1.merge(b2)
	if !b1.Equal(b3) {
		t.Fatal(b1, b3)
	}
}

func TestMerge1Reversed(t *testing.T) {
	var b1, b2, b3 *bkt

	// === panic capture
	// === panic expected !
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code should have panicked !?")
		}
	}()
	// === end of panic capture

	b1 = new(bkt).add(1.).add(2.)
	b2 = new(bkt).add(3.)
	b3 = new(bkt).add(2.).add(1.).add(3.)
	b2.merge(b1)
	if !b1.Equal(b3) {
		t.Fatal(b2, b3)
	}
}
