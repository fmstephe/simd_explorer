package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtps2pd_256.go -out ../asm_vcvtps2pd_256.s -stubs ../stub_vcvtps2pd_256.go -pkg cvtps2pd
func main() {
	TEXT("vcvtps2pd256", NOSPLIT, "func(vals *[4]float32, ret *[4]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	valsX := XMM()
	VMOVDQU(Mem{Base: vals}, valsX)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VCVTPS2PD(valsX, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
