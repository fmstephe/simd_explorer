package number

import (
	"encoding/binary"
	"math"
)

// Convenient collection of number<->bytes conversion functions
// all little endian
var endian = binary.LittleEndian

func ToFloat64(bytes []byte) float64 {
	return math.Float64frombits(ToUint64(bytes))
}

func ToFloat32(bytes []byte) float32 {
	return math.Float32frombits(ToUint32(bytes))
}

func ToFloat32Slice(bytes []byte) []float32 {
	ret := []float32{}
	for i := 0; i < len(bytes); i += 4 {
		ret = append(ret, ToFloat32(bytes[i:]))
	}

	return ret
}

func ToUint64(bytes []byte) uint64 {
	return endian.Uint64(bytes)
}

func ToInt64(bytes []byte) int64 {
	return int64(ToUint64(bytes))
}

func ToUint32(bytes []byte) uint32 {
	return endian.Uint32(bytes)
}

func ToUint32Slice(bytes []byte) []uint32 {
	ret := []uint32{}
	for i := 0; i < len(bytes); i += 4 {
		ret = append(ret, ToUint32(bytes[i:]))
	}

	return ret
}

func ToInt32(bytes []byte) int32 {
	return int32(ToUint32(bytes))
}

func ToInt32Slice(bytes []byte) []int32 {
	ret := []int32{}
	for i := 0; i < len(bytes); i += 4 {
		ret = append(ret, ToInt32(bytes[i:]))
	}

	return ret
}

func ToUint16(bytes []byte) uint16 {
	return endian.Uint16(bytes)
}

func ToInt16(bytes []byte) int16 {
	return int16(ToUint64(bytes))
}

func ToUint8(bytes []byte) uint8 {
	return bytes[0]
}

func ToInt8(bytes []byte) int8 {
	return int8(ToUint64(bytes))
}

func Float64ToBytes(val float64) []byte {
	bytes := make([]byte, 8)
	endian.PutUint64(bytes, math.Float64bits(val))
	return bytes
}

func Float32ToBytes(val float32) []byte {
	bytes := make([]byte, 4)
	endian.PutUint32(bytes, math.Float32bits(val))
	return bytes
}

func Float32SliceToBytes(vals []float32) []byte {
	bytes := []byte{}
	for _, val := range vals {
		bytes = endian.AppendUint32(bytes, math.Float32bits(val))
	}
	return bytes
}

func Uint64ToBytes(val uint64) []byte {
	bytes := make([]byte, 8)
	endian.PutUint64(bytes, uint64(val))
	return bytes
}

func Int64ToBytes(val int64) []byte {
	return Uint64ToBytes(uint64(val))
}

func Uint32ToBytes(val uint32) []byte {
	bytes := make([]byte, 4)
	endian.PutUint32(bytes, uint32(val))
	return bytes
}

func Uint32SliceToBytes(vals []uint32) []byte {
	bytes := []byte{}
	for _, val := range vals {
		bytes = endian.AppendUint32(bytes, val)
	}
	return bytes
}

func Int32ToBytes(val int32) []byte {
	return Uint32ToBytes(uint32(val))
}

func Uint16ToBytes(val uint16) []byte {
	bytes := make([]byte, 2)
	endian.PutUint16(bytes, uint16(val))
	return bytes
}

func Int16ToBytes(val int16) []byte {
	return Uint16ToBytes(uint16(val))
}

func Uint8ToBytes(val byte) []byte {
	return []byte{val}
}

func Int8ToBytes(val int8) []byte {
	return Uint8ToBytes(uint8(val))
}
