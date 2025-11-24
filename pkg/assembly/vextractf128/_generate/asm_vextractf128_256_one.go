package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vextractf128_256_one.go -out ../asm_vextractf128_256_one.s -stubs ../stub_vextractf128_256_one.go -pkg vextractf128
func main() {
	TEXT("vextractf128256One", NOSPLIT, "func(base *[8]float32, ret *[4]float32)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base YMM register")
	regY1 := YMM()
	VMOVDQA(Mem{Base: base}, regY1)

	Comment("Extract upper 128-bit lane (imm8=1) directly to memory")
	VEXTRACTF128(U8(0x01), regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
