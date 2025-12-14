package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtps2dq_128.go -out ../asm_vcvtps2dq_128.s -stubs ../stub_vcvtps2dq_128.go -pkg cvtps2dq
func main() {
	TEXT("vcvtps2dq128", NOSPLIT, "func(vals *[4]float32, ret *[4]int32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	valsX := XMM()
	VMOVDQU(Mem{Base: vals}, valsX)

	retX := XMM()

	Comment("Execute the instruction being demonstrated")
	VCVTPS2DQ(valsX, retX)

	Comment("Write results into return memory address")
	VMOVDQU(retX, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
