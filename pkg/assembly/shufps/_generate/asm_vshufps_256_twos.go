package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vshufps_256_twos.go -out ../asm_vshufps_256_twos.s -stubs ../stub_vshufps_256_twos.go -pkg shufps
func main() {
	TEXT("vshufps256Twos", NOSPLIT, "func(vals1, vals2, ret *[8]float32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into YMM register")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals1}, regY1)

	Comment("Load vals2 into YMM register")
	regY2 := YMM()
	VMOVDQA(Mem{Base: vals2}, regY2)

	Comment("VSHUFPS imm8=0xAA (1010_1010b) per 128-bit lane: dst = [a2,a2,b2,b2 | a6,a6,b6,b6]")
	VSHUFPS(U8(0xAA), regY2, regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Return from function")
	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()
	RET()

	// generate!
	Generate()
}
