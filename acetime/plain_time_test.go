package acetime

import (
	"testing"
)

func TestPlainTimeToSeconds(t *testing.T) {
	if PlainTimeToSeconds(0, 0, 0) != 0 {
		t.Fatalf(`PlainTimeToSeconds(0, 0, 0) should be 0`)
	}
	if PlainTimeToSeconds(0, 0, 1) != 1 {
		t.Fatalf(`PlainTimeToSeconds(0, 0, 1) should be 1`)
	}
	if PlainTimeToSeconds(0, 1, 0) != 60 {
		t.Fatalf(`PlainTimeToSeconds(0, 1, 0) should be 60`)
	}
	if PlainTimeToSeconds(1, 0, 0) != 3600 {
		t.Fatalf(`PlainTimeToSeconds(1, 0, 0) should be 3600`)
	}
	if PlainTimeToSeconds(1, 1, 1) != (3600 + 60 + 1) {
		t.Fatalf(`PlainTimeToSeconds(1, 1, 1) should be 3661`)
	}
}

func TestPlainTimeFromSeconds(t *testing.T) {
	h, m, s := PlainTimeFromSeconds(0)
	if h != 0 || m != 0 || s != 0 {
		t.Fatalf(`PlainTimeFromSeconds(0) should be (0, 0, 0)`)
	}
	h, m, s = PlainTimeFromSeconds(1)
	if h != 0 || m != 0 || s != 1 {
		t.Fatalf(`PlainTimeFromSeconds(0) should be (0, 0, 1)`)
	}
	h, m, s = PlainTimeFromSeconds(60)
	if h != 0 || m != 1 || s != 0 {
		t.Fatalf(`PlainTimeFromSeconds(0) should be (0, 1, 0)`)
	}
	h, m, s = PlainTimeFromSeconds(3600)
	if h != 1 || m != 0 || s != 0 {
		t.Fatalf(`PlainTimeFromSeconds(0) should be (1, 0, 0)`)
	}
	h, m, s = PlainTimeFromSeconds(3661)
	if h != 1 || m != 1 || s != 1 {
		t.Fatalf(`PlainTimeFromSeconds(0) should be (1, 1, 1)`)
	}
}
