package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmulhuw_128.go -out ../asm_vpmulhuw_128.s -stubs ../stub_vpmulhuw_128.go -pkg pmulhuw
func main() {
	TEXT("vpmulhuw128", NOSPLIT, "func(vals1, vals2, ret *[8]uint16)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals1}, regX1)

	Comment("Load vals2 into XMM register")
	regX2 := XMM()
	VMOVDQA(Mem{Base: vals2}, regX2)

	Comment("Multiply packed unsigned 16-bit integers with VEX encoding; keep high 16 bits")
	VPMULHUW(regX2, regX1, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
