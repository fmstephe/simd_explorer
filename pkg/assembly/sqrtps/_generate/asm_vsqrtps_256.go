package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vsqrtps_256.go -out ../asm_vsqrtps_256.s -stubs ../stub_vsqrtps_256.go -pkg sqrtps
func main() {
	TEXT("vsqrtps256", NOSPLIT, "func(vals *[8]float32, ret *[8]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals}, regY1)

	Comment("Compute square root of packed float32 values with VEX encoding: regY1 = sqrt(regY1)")
	VSQRTPS(regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
