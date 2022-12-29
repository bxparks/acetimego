package acetime

import (
	"testing"
)

func TestDateTupleCompare(t *testing.T) {
  a := DateTuple{2000, 1, 1, 0, suffixW}
  b := DateTuple{2000, 1, 1, 0, suffixW}
  if ! (dateTupleCompare(&a, &b) == 0) {
		t.Fatalf("(2000, 1, 1, 0, w) == (2000, 1, 1, 0, w)")
	}

  bb := DateTuple{2000, 1, 1, 0, suffixS}
  if dateTupleCompare(&a, &bb) != 0 {
		t.Fatalf("(2000, 1, 1, 0, s) == (2000, 1, 1, 0, w)")
	}

  c := DateTuple{2000, 1, 1, 1, suffixW}
  if ! (dateTupleCompare(&a, &c) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2000, 1, 1, 1, w)")
	}

  d := DateTuple{2000, 1, 2, 0, suffixW}
  if ! (dateTupleCompare(&a, &d) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2000, 1, 2, 0, w)")
	}

  e := DateTuple{2000, 2, 1, 0, suffixW}
  if ! (dateTupleCompare(&a, &e) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2000, 2, 1, 0, w)")
	}

  f := DateTuple{2001, 1, 1, 0, suffixW}
  if ! (dateTupleCompare(&a, &f) < 0) {
		t.Fatalf("(2000, 1, 1, 0, w) < (2001, 1, 1, 0, w)")
	}
}
