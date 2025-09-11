package ui

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

type valueConverter struct {
	bitsize int
	base    int
}

func newValueConverter(bitsize, base int) *valueConverter {
	mustValidBitsize(bitsize)
	return &valueConverter{
		bitsize: bitsize,
		base:    base,
	}
}

func (c *valueConverter) stringToBytes(txt string) []byte {
	val := c.stringToUint64(txt)
	endian := binary.LittleEndian
	switch c.bitsize {
	case 8:
		return []byte{byte(val)}
	case 16:
		return endian.AppendUint16(make([]byte, 0, 2), uint16(val))
	case 32:
		return endian.AppendUint32(make([]byte, 0, 4), uint32(val))
	case 64:
		return endian.AppendUint64(make([]byte, 0, 8), uint64(val))
	default:
		panic("unreachable")
	}
}

func (c *valueConverter) bytesToString(bytes []byte) string {
	val := uint64(0)
	endian := binary.LittleEndian
	switch c.bitsize {
	case 8:
		val = uint64(bytes[0])
	case 16:
		val = uint64(endian.Uint16(bytes))
	case 32:
		val = uint64(endian.Uint32(bytes))
	case 64:
		val = endian.Uint64(bytes)
	}

	return c.uint64ToString(val)
}

func (c *valueConverter) stringToUint64(txt string) uint64 {
	if txt == "" {
		// If the value of the field is empty default it to 0
		txt = c.uint64ToString(0)
	}
	val, err := strconv.ParseUint(txt, c.base, c.bitsize)
	if err != nil {
		panic(fmt.Errorf("Unexpected value %q found in register input, expecting decimal with bitsize %d: %s", txt, c.bitsize, err))
	}
	return val
}

func (c *valueConverter) uint64ToString(val uint64) string {
	return strconv.FormatUint(val, c.base)
}

// InputFieldInteger accepts unsigned integers.
func (c *valueConverter) inputAcceptor() func(string, rune) bool {
	base := c.base
	bitsize := c.bitsize
	return func(txt string, _ rune) bool {
		_, err := strconv.ParseUint(txt, base, bitsize)
		return err == nil
	}
}
