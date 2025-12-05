package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmulhw_128.go -out ../asm_vpmulhw_128.s -stubs ../stub_vpmulhw_128.go -pkg pmulhw
func main() {
	TEXT("vpmulhw128", NOSPLIT, "func(vals1 *[8]int16, vals2 *[8]int16, ret *[8]int16)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 and vals2 into XMM registers")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals1}, regX1)
	regX2 := XMM()
	VMOVDQA(Mem{Base: vals2}, regX2)

	Comment("Multiply packed signed 16-bit integers; keep high 16 bits of 32-bit products")
	VPMULHW(regX2, regX1, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
