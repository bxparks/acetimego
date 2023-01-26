package zoneinfo

import (
	"testing"
)

const data = "\x01\x01\x02\x01\x02\x03\x04"

func TestDataIO(t *testing.T) {
	f := NewDataIO(data)
	if f.err {
		t.Fatal(f.err)
	}

	a := f.ReadU8()
	if f.err != false {
		t.Fatal(f.err)
	}
	if a != 1 {
		t.Fatal(a)
	}

	b := f.ReadU16()
	if f.err != false {
		t.Fatal(f.err)
	}
	if b != 0x0201 {
		t.Fatal(b)
	}

	c := f.ReadU32()
	if f.err != false {
		t.Fatal(f.err)
	}
	if c != 0x04030201 {
		t.Fatal(c)
	}

	f.ReadU8()
	if f.err != true {
		t.Fatal(f.err)
	}
}
