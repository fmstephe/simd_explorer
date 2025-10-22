package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_sqrtss_128.go -out ../asm_sqrtss_128.s -stubs ../stub_sqrtss_128.go -pkg sqrtss
func main() {
	TEXT("sqrtss128", NOSPLIT, "func(vals *[4]float32, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Compute square root of scalar single precision value (lowest 32 bits)")
	SQRTSS(regX1, regX1) // Go assembler: regX1 = sqrt(regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}