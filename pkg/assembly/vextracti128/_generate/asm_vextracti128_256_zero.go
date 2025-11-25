package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vextracti128_256_zero.go -out ../asm_vextracti128_256_zero.s -stubs ../stub_vextracti128_256_zero.go -pkg vextracti128
func main() {
	TEXT("vextracti128256Zero", NOSPLIT, "func(vals256 *[8]uint32, ret *[4]uint32)")
	Comment("load params")
	vals256 := Load(Param("vals256"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals256 into YMM register")
	regY := YMM()
	VMOVDQU(Mem{Base: vals256}, regY)

	Comment("Extract lower 128-bit lane (0) from YMM to memory")
	VEXTRACTI128(U8(0x00), regY, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()
	RET()

	// generate!
	Generate()
}
