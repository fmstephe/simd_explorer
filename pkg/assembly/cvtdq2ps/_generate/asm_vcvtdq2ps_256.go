package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtdq2ps_256.go -out ../asm_vcvtdq2ps_256.s -stubs ../stub_vcvtdq2ps_256.go -pkg cvtdq2ps
func main() {
	TEXT("vcvtdq2ps256", NOSPLIT, "func(vals *[8]int32, ret *[8]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	valsY := YMM()
	VMOVDQU(Mem{Base: vals}, valsY)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VCVTDQ2PS(valsY, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
