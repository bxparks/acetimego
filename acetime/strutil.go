// Low-level conversion routines from uint8 and uint16 to ASCII strings.
// The goal is to eliminate the dependency to fmt.Sprintf() to save memory on
// TinyGo microcontrollers.

package acetime

import (
	"strings"
)

func WriteUint8Pad2(b *strings.Builder, n uint8, pad byte) {
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

func WriteUint16Pad4(b *strings.Builder, n uint16, pad byte) {
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
		WriteUint8Pad2(b, d32, pad)
		d10pad = '0'
	}
	WriteUint8Pad2(b, d10, d10pad)
}
