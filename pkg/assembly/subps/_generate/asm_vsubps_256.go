package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vsubps_256.go -out ../asm_vsubps_256.s -stubs ../stub_vsubps_256.go -pkg subps
func main() {
	TEXT("vsubps256", NOSPLIT, "func(vals1, vals2 *[8]float32, ret *[8]float32)")
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

	Comment("Subtract packed float32 values with VEX encoding: regY2 = regY1 - regY2")
	VSUBPS(regY2, regY1, regY2)

	Comment("Write results into return memory address")
	VMOVDQA(regY2, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
