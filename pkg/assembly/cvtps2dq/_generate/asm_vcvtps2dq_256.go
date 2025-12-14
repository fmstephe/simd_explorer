package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtps2dq_256.go -out ../asm_vcvtps2dq_256.s -stubs ../stub_vcvtps2dq_256.go -pkg cvtps2dq
func main() {
	TEXT("vcvtps2dq256", NOSPLIT, "func(vals *[8]float32, ret *[8]int32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	valsY := YMM()
	VMOVDQU(Mem{Base: vals}, valsY)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VCVTPS2DQ(valsY, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
