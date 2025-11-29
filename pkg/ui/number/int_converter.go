package number

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type IntConverter struct {
	bitWidth int
	base     int
}

func NewIntConverter(bitWidth, base int) *IntConverter {
	return &IntConverter{
		bitWidth: bitWidth,
		base:     base,
	}
}

func (c *IntConverter) GetBitWidth() int {
	return c.bitWidth
}

func (c *IntConverter) GetTextWidth() int {
	switch c.bitWidth {
	case 8:
		return len(strconv.FormatInt(math.MinInt8, c.base)) + 1
	case 16:
		return len(strconv.FormatInt(math.MinInt16, c.base)) + 1
	case 32:
		return len(strconv.FormatInt(math.MinInt32, c.base)) + 1
	case 64:
		return len(strconv.FormatInt(math.MinInt64, c.base)) + 1
	default:
		panic("unreachable")
	}
}

func (c *IntConverter) GetBase() int {
	return c.base
}

func (c *IntConverter) StringToBytes(txt string) []byte {
	val := c.mustStringToInt64(txt)
	switch c.bitWidth {
	case 8:
		return Int8ToBytes(int8(val))
	case 16:
		return Int16ToBytes(int16(val))
	case 32:
		return Int32ToBytes(int32(val))
	case 64:
		return Int64ToBytes(int64(val))
	default:
		panic("unreachable")
	}
}

func (c *IntConverter) BytesToString(bytes []byte) string {
	val := int64(0)
	switch c.bitWidth {
	case 8:
		val = int64(ToInt8(bytes))
	case 16:
		val = int64(ToInt16(bytes))
	case 32:
		val = int64(ToInt32(bytes))
	case 64:
		val = ToInt64(bytes)
	}

	return c.int64ToString(val)
}

func (c *IntConverter) Normalised(txt string) (normalised string, changed bool) {
	normalised = normaliseNegatives(txt)
	return normalised, normalised != txt
}

func (c *IntConverter) IsStable(_ string) bool {
	return true
}

// InputFieldInteger accepts unsigned integers.
func (c *IntConverter) InputAcceptor() func(string, rune) bool {
	return func(txt string, _ rune) bool {
		normalised, _ := c.Normalised(txt)
		_, err := c.stringToInt64(normalised)
		return err == nil
	}
}

func (c *IntConverter) mustStringToInt64(txt string) int64 {
	val, err := c.stringToInt64(txt)
	if err != nil {
		panic(fmt.Errorf("unexpected value %q found in register input, expecting signed integer with bitWidth %d: %s", txt, c.bitWidth, err))
	}
	return val
}

func (c *IntConverter) stringToInt64(txt string) (int64, error) {
	txt = strings.TrimSpace(txt)
	if txt == "" {
		// If the value of the field is empty default it to 0
		return 0, nil
	}
	if txt == "-" {
		// We allow a hanging '-', it's not a valid int but if we
		// don't parse it the user can't type any negative int value,
		// which necessarily _must_ start with a single '-'
		return 0, nil
	}
	return strconv.ParseInt(txt, c.base, c.bitWidth)
}

func (c *IntConverter) int64ToString(val int64) string {
	return strconv.FormatInt(val, c.base)
}
