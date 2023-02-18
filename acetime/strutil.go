// Low-level conversion routines from uint8 and uint16 to ASCII strings.
// The goal is to eliminate the dependency to fmt.Sprintf() to save memory on
// TinyGo microcontrollers.

package acetime

import (
	"strings"
)

// BuildUint8Pad2 converts the uint8 `n` into a decimal string of 2 spaces wide,
// padded on the left by the `pad` character.
func BuildUint8Pad2(b *strings.Builder, n uint8, pad byte) {
	if n >= 100 {
		b.WriteByte('*')
		b.WriteByte('*')
		return
	}
	d1 := n / 10
	d0 := n % 10
	var c1 byte
	if d1 == 0 {
		c1 = pad
	} else {
		c1 = d1 + '0'
	}
	c0 := d0 + '0'
	b.WriteByte(c1)
	b.WriteByte(c0)
}

// BuildUint8Pad2 converts uint16 `n` into a decimal string of 4-spaces wide,
// padded on the left by the `pad` character.
func BuildUint16Pad4(b *strings.Builder, n uint16, pad byte) {
	if n >= 10000 {
		b.WriteByte('*')
		b.WriteByte('*')
		b.WriteByte('*')
		b.WriteByte('*')
		return
	}

	d32 := uint8(n / 100)
	d10 := uint8(n % 100)
	var d10pad byte
	if d32 == 0 {
		b.WriteByte(pad)
		b.WriteByte(pad)
		d10pad = pad
	} else {
		BuildUint8Pad2(b, d32, pad)
		d10pad = '0'
	}
	BuildUint8Pad2(b, d10, d10pad)
}

// BuildUint64 converts uint64 `n` into a decimal string with no padding.
func BuildUint64(b *strings.Builder, n uint64) {
	// max uint64 is 1.8447e19, so 20 digits should be enough.
	var buf [20]uint8
	var i uint8
	var r uint8
	for i = 0; i < 20; i++ {
		r = uint8(n % 10)
		n = n / 10
		if n == 0 {
			if r != 0 || i == 0 {
				buf[i] = r
				i++
			}
			// i is the number of digits written
			break
		}
		buf[i] = r
	}

	// Write the digits backwards starting with buf[i-1]
	for ; i > 0; i-- {
		c := buf[i-1] + '0'
		b.WriteByte(c)
	}
}

// BuildTimeOffset converts the offsetSeconds into a string of the form
// +/-hh:mm, ignoring the remaining seconds component if any.
func BuildTimeOffset(b *strings.Builder, offsetSeconds int32) {
	sign, h, m, _ := secondsToHMS(offsetSeconds)
	var c byte
	if sign < 0 {
		c = '-'
	} else {
		c = '+'
	}

	b.WriteByte(c)
	BuildUint8Pad2(b, h, '0')
	b.WriteByte(':')
	BuildUint8Pad2(b, m, '0')
}

func secondsToHMS(seconds int32) (sign int8, h uint8, m uint8, s uint8) {
	if seconds < 0 {
		sign = -1
		seconds = -seconds
	} else {
		sign = 1
	}
	s = uint8(seconds % 60)
	minutes := seconds / 60
	m = uint8(minutes % 60)
	hours := uint8(minutes / 60)
	h = hours

	return
}
