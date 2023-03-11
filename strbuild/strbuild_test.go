package strbuild

import (
	"strings"
	"testing"
)

func TestUint8Pad2(t *testing.T) {
	var b strings.Builder

	b.Reset()
	Uint8Pad2(&b, 100, ' ')
	s := b.String()
	if !(s == "**") {
		t.Fatal(s)
	}

	b.Reset()
	Uint8Pad2(&b, 1, ' ')
	s = b.String()
	if !(s == " 1") {
		t.Fatal(s)
	}

	b.Reset()
	Uint8Pad2(&b, 1, '0')
	s = b.String()
	if !(s == "01") {
		t.Fatal(s)
	}

	b.Reset()
	Uint8Pad2(&b, 42, ' ')
	s = b.String()
	if !(s == "42") {
		t.Fatal(s)
	}
}

func TestUint16Pad4(t *testing.T) {
	var b strings.Builder

	b.Reset()
	Uint16Pad4(&b, 10000, ' ')
	s := b.String()
	if !(s == "****") {
		t.Fatal(s)
	}

	b.Reset()
	Uint16Pad4(&b, 1, ' ')
	s = b.String()
	if !(s == "   1") {
		t.Fatal(s)
	}

	b.Reset()
	Uint16Pad4(&b, 1, '0')
	s = b.String()
	if !(s == "0001") {
		t.Fatal(s)
	}

	b.Reset()
	Uint16Pad4(&b, 42, ' ')
	s = b.String()
	if !(s == "  42") {
		t.Fatal(s)
	}

	b.Reset()
	Uint16Pad4(&b, 42, '0')
	s = b.String()
	if !(s == "0042") {
		t.Fatal(s)
	}

	b.Reset()
	Uint16Pad4(&b, 421, ' ')
	s = b.String()
	if !(s == " 421") {
		t.Fatal(s)
	}

	b.Reset()
	Uint16Pad4(&b, 421, '0')
	s = b.String()
	if !(s == "0421") {
		t.Fatal(s)
	}
}

func TestUint64(t *testing.T) {
	var b strings.Builder

	Uint64(&b, 0)
	s := b.String()
	if !(s == "0") {
		t.Fatal(s)
	}

	b.Reset()
	Uint64(&b, 1)
	s = b.String()
	if !(s == "1") {
		t.Fatal(s)
	}

	b.Reset()
	Uint64(&b, 10)
	s = b.String()
	if !(s == "10") {
		t.Fatal(s)
	}

	b.Reset()
	Uint64(&b, 12)
	s = b.String()
	if !(s == "12") {
		t.Fatal(s)
	}

	b.Reset()
	Uint64(&b, 1234)
	s = b.String()
	if !(s == "1234") {
		t.Fatal(s)
	}

	b.Reset()
	Uint64(&b, 12345678)
	s = b.String()
	if !(s == "12345678") {
		t.Fatal(s)
	}

	b.Reset()
	Uint64(&b, 1234567890123456)
	s = b.String()
	if !(s == "1234567890123456") {
		t.Fatal(s)
	}

	b.Reset()
	Uint64(&b, 1234567890123456789)
	s = b.String()
	if !(s == "1234567890123456789") {
		t.Fatal(s)
	}
}

func TestTimeOffset(t *testing.T) {
	var b strings.Builder

	// 1h2m
	TimeOffset(&b, 1*3600+2*60)
	s := b.String()
	if !(s == "+01:02") {
		t.Fatal(s)
	}

	// Second component ignored.
	b.Reset()
	TimeOffset(&b, 1*3600+2*60+1)
	s = b.String()
	if !(s == "+01:02") {
		t.Fatal(s)
	}

	// Negative -1h2m
	b.Reset()
	TimeOffset(&b, -1*3600+-2*60)
	s = b.String()
	if !(s == "-01:02") {
		t.Fatal(s)
	}
}

func TestSecondsToHMS(t *testing.T) {
	sign, h, m, s := secondsToHMS(0)
	if !(sign == 1) {
		t.Fatal(sign)
	}
	if !(h == 0) {
		t.Fatal(h)
	}
	if !(m == 0) {
		t.Fatal(m)
	}
	if !(s == 0) {
		t.Fatal(s)
	}

	sign, h, m, s = secondsToHMS(1)
	if !(sign == 1) {
		t.Fatal(sign)
	}
	if !(h == 0) {
		t.Fatal(h)
	}
	if !(m == 0) {
		t.Fatal(m)
	}
	if !(s == 1) {
		t.Fatal(s)
	}

	sign, h, m, s = secondsToHMS(62)
	if !(sign == 1) {
		t.Fatal(sign)
	}
	if !(h == 0) {
		t.Fatal(h)
	}
	if !(m == 1) {
		t.Fatal(m)
	}
	if !(s == 2) {
		t.Fatal(s)
	}

	sign, h, m, s = secondsToHMS(-3663)
	if !(sign == -1) {
		t.Fatal(sign)
	}
	if !(h == 1) {
		t.Fatal(h)
	}
	if !(m == 1) {
		t.Fatal(m)
	}
	if !(s == 3) {
		t.Fatal(s)
	}
}
