package register

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
)

type valueConverter struct {
	bitsize   int
	base      int
	textWidth int
}

func newValueConverter(bitsize, base int) *valueConverter {
	mustValidPartSize(bitsize)
	return &valueConverter{
		bitsize:   bitsize,
		base:      base,
		textWidth: calculateTextWidth(bitsize, base),
	}
}

func (c *valueConverter) stringToBytes(txt string) []byte {
	val := c.stringToUint64(txt)
	switch c.bitsize {
	case 8:
		return asmutil.Uint8ToBytes(uint8(val))
	case 16:
		return asmutil.Uint16ToBytes(uint16(val))
	case 32:
		return asmutil.Uint32ToBytes(uint32(val))
	case 64:
		return asmutil.Uint64ToBytes(uint64(val))
	default:
		panic("unreachable")
	}
}

func (c *valueConverter) bytesToString(bytes []byte) string {
	val := uint64(0)
	switch c.bitsize {
	case 8:
		val = uint64(asmutil.ToUint8(bytes))
	case 16:
		val = uint64(asmutil.ToUint16(bytes))
	case 32:
		val = uint64(asmutil.ToUint32(bytes))
	case 64:
		val = asmutil.ToUint64(bytes)
	}

	return c.uint64ToString(val)
}

func (c *valueConverter) stringToUint64(txt string) uint64 {
	txt = strings.TrimSpace(txt)
	if txt == "" {
		// If the value of the field is empty default it to 0
		return 0
	}
	val, err := strconv.ParseUint(txt, c.base, c.bitsize)
	if err != nil {
		panic(fmt.Errorf("Unexpected value %q found in register input, expecting decimal with bitsize %d: %s", txt, c.bitsize, err))
	}
	return val
}

func (c *valueConverter) uint64ToString(val uint64) string {
	raw := strconv.FormatUint(val, c.base)
	return c.leftPad(raw)
}

func (c *valueConverter) leftPad(txt string) string {
	if len(txt) > c.textWidth {
		panic(fmt.Errorf("Attempted to process string too long (%d) for bitsize (%d) and base (%d) string must be %d or shorter", len(txt), c.bitsize, c.base, c.textWidth))
	}
	return strings.Repeat("0", (c.textWidth-1)-len(txt)) + txt
}

// InputFieldInteger accepts unsigned integers.
func (c *valueConverter) inputAcceptor() func(string, rune) bool {
	base := c.base
	bitsize := c.bitsize
	return func(txt string, _ rune) bool {
		txt = strings.TrimSpace(txt)
		_, err := strconv.ParseUint(txt, base, bitsize)
		return err == nil
	}
}

func (c *valueConverter) getTextWidth() int {
	return c.textWidth
}

func calculateTextWidth(bitsize, base int) int {
	mustValidPartSize(bitsize)
	switch bitsize {
	case 8:
		return len(strconv.FormatUint(math.MaxUint8, base)) + 1
	case 16:
		return len(strconv.FormatUint(math.MaxUint16, base)) + 1
	case 32:
		return len(strconv.FormatUint(math.MaxUint32, base)) + 1
	case 64:
		return len(strconv.FormatUint(math.MaxUint64, base)) + 1
	default:
		panic("unreachable")
	}
}
