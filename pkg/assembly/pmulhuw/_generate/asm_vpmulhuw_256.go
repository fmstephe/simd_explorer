package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmulhuw_256.go -out ../asm_vpmulhuw_256.s -stubs ../stub_vpmulhuw_256.go -pkg pmulhuw
func main() {
	TEXT("vpmulhuw256", NOSPLIT, "func(vals1, vals2, ret *[16]uint16)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into YMM register")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals1}, regY1)

	Comment("Load vals2 into YMM register")
	regY2 := YMM()
	VMOVDQA(Mem{Base: vals2}, regY2)

	Comment("Multiply packed unsigned 16-bit integers (per 128-bit lane); keep high 16 bits")
	VPMULHUW(regY2, regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
