package zoneinfo

// StringIO8 holds multiple strings concatenated into a single string, and
// provides the StringAt() function to extract a given substring. The offsets
// must have a terminating sentinel entry which is just past the data string.
// In other words, len(offsets)==numStrings+1.
// This type uses uint8 offsets, which limits the maximum concatenated string to
// 254 (2^8-2) characters. (The character at index 255 is inaccessible).
type StringIO8 struct {
	data    string  // concatenated string
	offsets []uint8 // has terminating sentinel
}

func (d *StringIO8) StringAt(i uint8) string {
	begin := d.offsets[i]
	end := d.offsets[i+1] // always exists because of terminating sentinel
	return d.data[begin:end]
}

// StringIO16 holds multiple strings concatenated into a single string, and
// provides the StringAt() function to extract a given substring. The offsets
// must have a terminating sentinel entry which is just past the data string.
// In other words, len(offsets)==numStrings+1. This type uses uint16 offsets,
// which limits the maximum concatenated string to 65534 (2^16-2) characters.
// (The character at index 65535 is inaccessible).
type StringIO16 struct {
	data    string   // concatenated string
	offsets []uint16 // has terminating sentinel
}

func (d *StringIO16) StringAt(i uint16) string {
	begin := d.offsets[i]
	end := d.offsets[i+1] // always exists because of terminating sentinel
	return d.data[begin:end]
}
