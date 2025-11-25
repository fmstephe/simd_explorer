package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vunpckhps_256.go -out ../asm_vunpckhps_256.s -stubs ../stub_vunpckhps_256.go -pkg unpckhps
func main() {
	TEXT("vunpckhps256", NOSPLIT, "func(vals1, vals2, ret *[8]float32)")
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

	Comment("VUNPCKHPS per 128-bit lane: dst = [a2,b2,a3,b3 | a6,b6,a7,b7]")
	VUNPCKHPS(regY2, regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()
	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
