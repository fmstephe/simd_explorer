package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvttss2si_128_int32.go -out ../asm_vcvttss2si_128_int32.s -stubs ../stub_vcvttss2si_128_int32.go -pkg cvttss2si
func main() {
	TEXT("vcvttss2si128Int32", NOSPLIT, "func(vals *[4]float32, ret *int32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Truncate scalar single (lowest lane) to signed 32-bit integer")
	reg32 := GP32()
	VCVTTSS2SI(regX1, reg32)

	Comment("Write result to return memory address")
	MOVL(reg32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
