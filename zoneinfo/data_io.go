package zoneinfo

// Simple I/O interface to extract integer fields of various sizes from a string
// containing arbitrary sequence of 8-bit binary data. A string is immutable and
// is guaranteed to be stored in flash memory instead of being copied into RAM.
// The ZoneInfos, ZoneEras, ZonePolicies, ZoneEras are encoded into binary and
// stored as strings. This class is able to extract the binary data and
// rehydrate it.

type DataIO struct {
	data string // original binary data, no modified
	s    string // current slice into data
	err  bool
}

func NewDataIO(data string) DataIO {
	return DataIO{
		data: data,
		s:    data,
		err:  false,
	}
}

func (f *DataIO) Seek(n uint16) {
	f.s = f.data[n:]
}

func (f *DataIO) Reset() {
	f.s = f.data
}

// read n bytes as a string
func (f *DataIO) read(n uint8) string {
	if len(f.s) < int(n) {
		f.err = true
		return ""
	}
	s := f.s[0:n]
	f.s = f.s[n:]
	return s
}

// ReadU8() reads in 1 byte as uint8.
func (f *DataIO) ReadU8() uint8 {
	s := f.read(1)
	if s == "" {
		return 0
	}
	return uint8(s[0])
}

// ReadU16() reads in 2 bytes in little-endian format as uint16.
func (f *DataIO) ReadU16() uint16 {
	s := f.read(2)
	if s == "" {
		return 0
	}
	return uint16(s[0]) | uint16(s[1])<<8
}

// ReadU32() reads in 4 bytes in little-endian format as uint32
func (f *DataIO) ReadU32() uint32 {
	s := f.read(4)
	if s == "" {
		return 0
	}
	return uint32(s[0]) | uint32(s[1])<<8 | uint32(s[2])<<16 | uint32(s[3])<<24
}
