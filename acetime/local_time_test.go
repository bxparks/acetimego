package acetime

import (
	"testing"
)

func TestLocalTimeToSeconds(t *testing.T) {
	if LocalTimeToSeconds(0, 0, 0) != 0 {
		t.Fatalf(`LocalTimeToSeconds(0, 0, 0) should be 0`)
	}
	if LocalTimeToSeconds(0, 0, 1) != 1 {
		t.Fatalf(`LocalTimeToSeconds(0, 0, 1) should be 1`)
	}
	if LocalTimeToSeconds(0, 1, 0) != 60 {
		t.Fatalf(`LocalTimeToSeconds(0, 1, 0) should be 60`)
	}
	if LocalTimeToSeconds(1, 0, 0) != 3600 {
		t.Fatalf(`LocalTimeToSeconds(1, 0, 0) should be 3600`)
	}
	if LocalTimeToSeconds(1, 1, 1) != (3600 + 60 + 1) {
		t.Fatalf(`LocalTimeToSeconds(1, 1, 1) should be 3661`)
	}
}

func TestLocalTimeFromSeconds(t *testing.T) {
	h, m, s := LocalTimeFromSeconds(0)
	if h != 0 || m != 0 || s != 0 {
		t.Fatalf(`LocalTimeFromSeconds(0) should be (0, 0, 0)`)
	}
	h, m, s = LocalTimeFromSeconds(1)
	if h != 0 || m != 0 || s != 1 {
		t.Fatalf(`LocalTimeFromSeconds(0) should be (0, 0, 1)`)
	}
	h, m, s = LocalTimeFromSeconds(60)
	if h != 0 || m != 1 || s != 0 {
		t.Fatalf(`LocalTimeFromSeconds(0) should be (0, 1, 0)`)
	}
	h, m, s = LocalTimeFromSeconds(3600)
	if h != 1 || m != 0 || s != 0 {
		t.Fatalf(`LocalTimeFromSeconds(0) should be (1, 0, 0)`)
	}
	h, m, s = LocalTimeFromSeconds(3661)
	if h != 1 || m != 1 || s != 1 {
		t.Fatalf(`LocalTimeFromSeconds(0) should be (1, 1, 1)`)
	}
}
