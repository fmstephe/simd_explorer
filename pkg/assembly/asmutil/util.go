package asmutil

import "encoding/binary"

var endian = binary.LittleEndian

func ToUint64(bytes []byte) uint64 {
	return endian.Uint64(bytes)
}

func ToUint32(bytes []byte) uint32 {
	return endian.Uint32(bytes)
}

func ToUint16(bytes []byte) uint16 {
	return endian.Uint16(bytes)
}

func ToUint8(bytes []byte) uint8 {
	return bytes[0]
}

func Uint64ToBytes(val uint64) []byte {
	bytes := make([]byte, 8)
	endian.PutUint64(bytes, uint64(val))
	return bytes
}

func Uint32ToBytes(val uint32) []byte {
	bytes := make([]byte, 4)
	endian.PutUint32(bytes, uint32(val))
	return bytes
}

func Uint16ToBytes(val uint16) []byte {
	bytes := make([]byte, 2)
	endian.PutUint16(bytes, uint16(val))
	return bytes
}

func Uint8ToBytes(val byte) []byte {
	return []byte{val}
}
