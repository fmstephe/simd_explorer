package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtsi2ss_128_int32.go -out ../asm_vcvtsi2ss_128_int32.s -stubs ../stub_vcvtsi2ss_128_int32.go -pkg cvtsi2ss
func main() {
	TEXT("vcvtsi2ss128int32", NOSPLIT, "func(vals *[4]float32, ival *int32, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	scalar := Load(Param("ival"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Convert signed 32-bit integer to scalar single and insert into lowest lane")
	VCVTSI2SSL(Mem{Base: scalar}, regX1, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
