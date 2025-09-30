package number

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type UintConverter struct {
	bitWidth int
	base     int
}

func NewUintConverter(bitWidth, base int) *UintConverter {
	return &UintConverter{
		bitWidth: bitWidth,
		base:     base,
	}
}

func (c *UintConverter) GetBitWidth() int {
	return c.bitWidth
}

func (c *UintConverter) GetTextWidth() int {
	switch c.bitWidth {
	case 8:
		return len(strconv.FormatUint(math.MaxUint8, c.base)) + 1
	case 16:
		return len(strconv.FormatUint(math.MaxUint16, c.base)) + 1
	case 32:
		return len(strconv.FormatUint(math.MaxUint32, c.base)) + 1
	case 64:
		return len(strconv.FormatUint(math.MaxUint64, c.base)) + 1
	default:
		panic("unreachable")
	}
}

func (c *UintConverter) GetBase() int {
	return c.base
}

func (c *UintConverter) StringToBytes(txt string) []byte {
	val := c.stringToUint64(txt)
	switch c.bitWidth {
	case 8:
		return Uint8ToBytes(uint8(val))
	case 16:
		return Uint16ToBytes(uint16(val))
	case 32:
		return Uint32ToBytes(uint32(val))
	case 64:
		return Uint64ToBytes(uint64(val))
	default:
		panic("unreachable")
	}
}

func (c *UintConverter) BytesToString(bytes []byte) string {
	val := uint64(0)
	switch c.bitWidth {
	case 8:
		val = uint64(ToUint8(bytes))
	case 16:
		val = uint64(ToUint16(bytes))
	case 32:
		val = uint64(ToUint32(bytes))
	case 64:
		val = ToUint64(bytes)
	}

	return c.uint64ToString(val)
}

// InputFieldInteger accepts unsigned integers.
func (c *UintConverter) InputAcceptor() func(string, rune) bool {
	return func(txt string, _ rune) bool {
		_, err := c.stringToUint64Err(txt)
		return err == nil
	}
}

func (c *UintConverter) stringToUint64(txt string) uint64 {
	val, err := c.stringToUint64Err(txt)
	if err != nil {
		panic(fmt.Errorf("Unexpected value %q found in register input, expecting unsigned integer with bitWidth %d: %s", txt, c.bitWidth, err))
	}
	return val
}

func (c *UintConverter) stringToUint64Err(txt string) (uint64, error) {
	txt = strings.TrimSpace(txt)
	if txt == "" {
		// If the value of the field is empty default it to 0
		return 0, nil
	}
	return strconv.ParseUint(txt, c.base, c.bitWidth)
}

func (c *UintConverter) uint64ToString(val uint64) string {
	raw := strconv.FormatUint(val, c.base)
	return c.leftPad(raw)
}

func (c *UintConverter) leftPad(txt string) string {
	if len(txt) > c.GetTextWidth() {
		panic(fmt.Errorf("Attempted to process string too long (%d) for bitWidth (%d) and base (%d) string must be %d or shorter", len(txt), c.bitWidth, c.base, c.GetTextWidth()))
	}
	return strings.Repeat("0", (c.GetTextWidth()-1)-len(txt)) + txt
}
