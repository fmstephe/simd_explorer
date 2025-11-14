package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vshufps_128_threes.go -out ../asm_vshufps_128_threes.s -stubs ../stub_vshufps_128_threes.go -pkg shufps
func main() {
	TEXT("vshufps128Threes", NOSPLIT, "func(vals1, vals2, ret *[4]float32)")
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

	Comment("VSHUFPS imm8=0xFF (1111_1111b): dst = [a3,a3,b3,b3]")
	VSHUFPS(U8(0xFF), regX2, regX1, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
