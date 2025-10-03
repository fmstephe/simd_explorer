package number

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type FloatConverter struct {
	bitWidth int
}

func NewFloatConverter(bitWidth int) *FloatConverter {
	return &FloatConverter{
		bitWidth: bitWidth,
	}
}

func (c *FloatConverter) GetBitWidth() int {
	return c.bitWidth
}

func (c *FloatConverter) GetTextWidth() int {
	switch c.bitWidth {
	case 32:
		return len(strconv.FormatFloat(math.MaxFloat32, 'f', -1, c.bitWidth)) + 1
	case 64:
		return len(strconv.FormatFloat(math.MaxFloat64, 'f', -1, c.bitWidth)) + 1
	default:
		panic("unreachable")
	}
}

func (c *FloatConverter) GetBase() int {
	return 10
}

func (c *FloatConverter) StringToBytes(txt string) []byte {
	val := c.mustStringToFloat64(txt)
	switch c.bitWidth {
	case 32:
		return Float32ToBytes(float32(val))
	case 64:
		return Float64ToBytes(float64(val))
	default:
		panic("unreachable")
	}
}

func (c *FloatConverter) BytesToString(bytes []byte) string {
	val := float64(0)
	switch c.bitWidth {
	case 32:
		val = float64(ToFloat32(bytes))
	case 64:
		val = ToFloat64(bytes)
	}

	return c.float64ToString(val)
}

func (c *FloatConverter) IsStable(txt string) bool {
	f := c.mustStringToFloat64(txt)
	txt2 := c.float64ToString(f)
	f2 := c.mustStringToFloat64(txt2)

	return txt == txt2 && f == f2
}

// InputFieldInteger accepts unsigned integers.
func (c *FloatConverter) InputAcceptor() func(string, rune) bool {
	return func(txt string, _ rune) bool {
		_, err := c.stringToFloat64(txt)
		return err == nil
	}
}

func (c *FloatConverter) mustStringToFloat64(txt string) float64 {
	val, err := c.stringToFloat64(txt)
	if err != nil {
		panic(fmt.Errorf("Unexpected value %q found in register input, expecting unsigned integer with bitWidth %d: %s", txt, c.bitWidth, err))
	}
	return val
}

func (c *FloatConverter) stringToFloat64(txt string) (float64, error) {
	txt = strings.TrimSpace(txt)
	if txt == "" {
		// If the value of the field is empty default it to 0
		return 0, nil
	}
	return strconv.ParseFloat(txt, c.bitWidth)
}

func (c *FloatConverter) float64ToString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, c.bitWidth)
}
