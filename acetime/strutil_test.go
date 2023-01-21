package acetime

import (
	"testing"
	"strings"
)

func TestWriteUint8Pad2(t *testing.T) {
	var b strings.Builder

	b.Reset()
	WriteUint8Pad2(&b, 100, ' ')
	s := b.String()
	if !(s == "**") {
		t.Fatal(s)
	}

	b.Reset()
	WriteUint8Pad2(&b, 1, ' ')
	s = b.String()
	if !(s == " 1") {
		t.Fatal(s)
	}

	b.Reset()
	WriteUint8Pad2(&b, 1, '0')
	s = b.String()
	if !(s == "01") {
		t.Fatal(s)
	}

	b.Reset()
	WriteUint8Pad2(&b, 42, ' ')
	s = b.String()
	if !(s == "42") {
		t.Fatal(s)
	}
}

func TestWriteUint16Pad4(t *testing.T) {
	var b strings.Builder

	b.Reset()
	WriteUint16Pad4(&b, 10000, ' ')
	s := b.String()
	if !(s == "****") {
		t.Fatal(s)
	}

	b.Reset()
	WriteUint16Pad4(&b, 1, ' ')
	s = b.String()
	if !(s == "   1") {
		t.Fatal(s)
	}

	b.Reset()
	WriteUint16Pad4(&b, 1, '0')
	s = b.String()
	if !(s == "0001") {
		t.Fatal(s)
	}

	b.Reset()
	WriteUint16Pad4(&b, 42, ' ')
	s = b.String()
	if !(s == "  42") {
		t.Fatal(s)
	}

	b.Reset()
	WriteUint16Pad4(&b, 42, '0')
	s = b.String()
	if !(s == "0042") {
		t.Fatal(s)
	}

	b.Reset()
	WriteUint16Pad4(&b, 421, ' ')
	s = b.String()
	if !(s == " 421") {
		t.Fatal(s)
	}

	b.Reset()
	WriteUint16Pad4(&b, 421, '0')
	s = b.String()
	if !(s == "0421") {
		t.Fatal(s)
	}
}
