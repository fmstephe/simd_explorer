package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vrsqrtps_256.go -out ../asm_vrsqrtps_256.s -stubs ../stub_vrsqrtps_256.go -pkg rsqrtps
func main() {
	TEXT("vrsqrtps256", NOSPLIT, "func(vals *[8]float32, ret *[8]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals}, regY1)

	Comment("Compute reciprocal square root of packed float32 values with VEX encoding: regY1 = 1.0 / sqrt(regY1)")
	VRSQRTPS(regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
