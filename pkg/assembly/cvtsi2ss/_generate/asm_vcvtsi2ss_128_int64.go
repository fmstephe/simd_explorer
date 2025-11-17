package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vcvtsi2ss_128_int64.go -out ../asm_vcvtsi2ss_128_int64.s -stubs ../stub_vcvtsi2ss_128_int64.go -pkg cvtsi2ss
func main() {
	TEXT("vcvtsi2ss128int64", NOSPLIT, "func(vals *[4]float32, ival *int64, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	scalar := Load(Param("ival"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Convert signed 64-bit integer to scalar single and insert into lowest lane")
	VCVTSI2SSQ(Mem{Base: scalar}, regX1, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
