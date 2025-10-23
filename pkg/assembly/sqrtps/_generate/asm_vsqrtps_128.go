package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vsqrtps_128.go -out ../asm_vsqrtps_128.s -stubs ../stub_vsqrtps_128.go -pkg sqrtps
func main() {
	TEXT("vsqrtps128", NOSPLIT, "func(vals *[4]float32, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Compute square root of packed float32 values with VEX encoding: regX1 = sqrt(regX1)")
	VSQRTPS(regX1, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
