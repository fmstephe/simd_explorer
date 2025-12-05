package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmulhw_256.go -out ../asm_vpmulhw_256.s -stubs ../stub_vpmulhw_256.go -pkg pmulhw
func main() {
	TEXT("vpmulhw256", NOSPLIT, "func(vals1 *[16]int16, vals2 *[16]int16, ret *[16]int16)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 and vals2 into YMM registers")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals1}, regY1)
	regY2 := YMM()
	VMOVDQA(Mem{Base: vals2}, regY2)

	Comment("Multiply packed signed 16-bit integers; keep high 16 bits of 32-bit products")
	VPMULHW(regY2, regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
