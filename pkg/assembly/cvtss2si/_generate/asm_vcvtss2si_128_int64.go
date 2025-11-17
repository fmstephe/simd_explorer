package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtss2si_128_int64.go -out ../asm_vcvtss2si_128_int64.s -stubs ../stub_vcvtss2si_128_int64.go -pkg cvtss2si
func main() {
	TEXT("vcvtss2si128Int64", NOSPLIT, "func(vals *[4]float32, ret *int64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Convert scalar single in lowest lane to signed 64-bit integer")
	reg64 := GP64()
	VCVTSS2SIQ(regX1, reg64)

	Comment("Write result to return memory address")
	MOVQ(reg64, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
