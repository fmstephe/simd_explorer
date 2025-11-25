package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vshufps_256_ones.go -out ../asm_vshufps_256_ones.s -stubs ../stub_vshufps_256_ones.go -pkg shufps
func main() {
	TEXT("vshufps256Ones", NOSPLIT, "func(vals1, vals2, ret *[8]float32)")
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

	Comment("VSHUFPS imm8=0x55 (0101_0101b) per 128-bit lane: dst = [a1,a1,b1,b1 | a5,a5,b5,b5]")
	VSHUFPS(U8(0x55), regY2, regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Return from function")
	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()
	RET()

	// generate!
	Generate()
}
