package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_unpcklps_128.go -out ../asm_unpcklps_128.s -stubs ../stub_unpcklps_128.go -pkg unpcklps
func main() {
	TEXT("unpcklps128", NOSPLIT, "func(vals1, vals2, ret *[4]float32)")
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

	Comment("UNPCKLPS: interleave low pairs -> dst = [a0,b0,a1,b1]")
	// UNPCKLPS xmm1,xmm2 interleaves low 64 bits of xmm1 and xmm2 into xmm1
	UNPCKLPS(regX2, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
