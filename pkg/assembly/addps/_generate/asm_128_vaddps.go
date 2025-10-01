package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_128_vaddps.go -out ../asm_128_vaddps.s -stubs ../stub_128_vaddps.go -pkg addps
func main() {
	TEXT("vaddps128", NOSPLIT, "func(vals1, vals2, ret *[4]float32)")

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

	Comment("Sum packed float32 values from vals1 and vals2")
	VADDPS(regX2, regX1, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
