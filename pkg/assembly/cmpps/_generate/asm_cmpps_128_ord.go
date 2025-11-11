package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_cmpps_128_ord.go -out ../asm_cmpps_128_ord.s -stubs ../stub_cmpps_128_ord.go -pkg cmpps
func main() {
	TEXT("cmpps128Ord", NOSPLIT, "func(vals1, vals2, ret *[4]float32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals1}, regX1)

	Comment("Load vals2 into XMM register")
	regX2 := XMM()
	VMOVDQA(Mem{Base: vals2}, regX2)

	Comment("Compare packed single-precision: ORD (imm8=7)")
	CMPPS(regX2, regX1, U8(7))

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
