package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtdq2pd_256.go -out ../asm_vcvtdq2pd_256.s -stubs ../stub_vcvtdq2pd_256.go -pkg cvtdq2pd
func main() {
	TEXT("vcvtdq2pd256", NOSPLIT, "func(vals *[4]int32, ret *[4]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	valsX := XMM()
	VMOVDQU(Mem{Base: vals}, valsX)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VCVTDQ2PD(valsX, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
