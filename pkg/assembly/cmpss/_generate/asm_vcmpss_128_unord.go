package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcmpss_128_unord.go -out ../asm_vcmpss_128_unord.s -stubs ../stub_vcmpss_128_unord.go -pkg cmpss
func main() {
	TEXT("vcmpss128Unord", NOSPLIT, "func(vals1, vals2, ret *[4]float32)")
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

	Comment("Compare scalar single precision values for unordered (lowest 32 bits)")
	regX3 := XMM()
	VCMPSS(U8(3), regX2, regX1, regX3) // 3 = UNORD

	Comment("Write results into return memory address")
	VMOVDQA(regX3, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
