package acetime

import (
	"strings"
	"testing"
)

func TestBuildUint8Pad2(t *testing.T) {
	var b strings.Builder

	b.Reset()
	BuildUint8Pad2(&b, 100, ' ')
	s := b.String()
	if !(s == "**") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint8Pad2(&b, 1, ' ')
	s = b.String()
	if !(s == " 1") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint8Pad2(&b, 1, '0')
	s = b.String()
	if !(s == "01") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint8Pad2(&b, 42, ' ')
	s = b.String()
	if !(s == "42") {
		t.Fatal(s)
	}
}

func TestBuildUint16Pad4(t *testing.T) {
	var b strings.Builder

	b.Reset()
	BuildUint16Pad4(&b, 10000, ' ')
	s := b.String()
	if !(s == "****") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint16Pad4(&b, 1, ' ')
	s = b.String()
	if !(s == "   1") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint16Pad4(&b, 1, '0')
	s = b.String()
	if !(s == "0001") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint16Pad4(&b, 42, ' ')
	s = b.String()
	if !(s == "  42") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint16Pad4(&b, 42, '0')
	s = b.String()
	if !(s == "0042") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint16Pad4(&b, 421, ' ')
	s = b.String()
	if !(s == " 421") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint16Pad4(&b, 421, '0')
	s = b.String()
	if !(s == "0421") {
		t.Fatal(s)
	}
}

func TestBuildUint64(t *testing.T) {
	var b strings.Builder

	BuildUint64(&b, 0)
	s := b.String()
	if !(s == "0") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint64(&b, 1)
	s = b.String()
	if !(s == "1") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint64(&b, 10)
	s = b.String()
	if !(s == "10") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint64(&b, 12)
	s = b.String()
	if !(s == "12") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint64(&b, 1234)
	s = b.String()
	if !(s == "1234") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint64(&b, 12345678)
	s = b.String()
	if !(s == "12345678") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint64(&b, 1234567890123456)
	s = b.String()
	if !(s == "1234567890123456") {
		t.Fatal(s)
	}

	b.Reset()
	BuildUint64(&b, 1234567890123456789)
	s = b.String()
	if !(s == "1234567890123456789") {
		t.Fatal(s)
	}
}
