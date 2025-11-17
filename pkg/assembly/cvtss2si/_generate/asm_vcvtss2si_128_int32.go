package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtss2si_128_int32.go -out ../asm_vcvtss2si_128_int32.s -stubs ../stub_vcvtss2si_128_int32.go -pkg cvtss2si
func main() {
	TEXT("vcvtss2si128Int32", NOSPLIT, "func(vals *[4]float32, ret *int32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Convert scalar single in lowest lane to signed 32-bit integer")
	reg32 := GP32()
	VCVTSS2SI(regX1, reg32)

	Comment("Write result to return memory address")
	MOVL(reg32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
