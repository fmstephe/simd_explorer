package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vextractf128_256_zero.go -out ../asm_vextractf128_256_zero.s -stubs ../stub_vextractf128_256_zero.go -pkg vextractf128
func main() {
	TEXT("vextractf128256Zero", NOSPLIT, "func(base *[8]float32, ret *[4]float32)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base YMM register")
	regY1 := YMM()
	VMOVDQA(Mem{Base: base}, regY1)

	Comment("Extract lower 128-bit lane (imm8=0) directly to memory")
	VEXTRACTF128(U8(0x00), regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
