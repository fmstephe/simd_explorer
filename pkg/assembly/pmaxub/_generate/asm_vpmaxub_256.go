package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaxub_256.go -out ../asm_vpmaxub_256.s -stubs ../stub_vpmaxub_256.go -pkg pmaxub
func main() {
	TEXT("vpmaxub256", NOSPLIT, "func(vals1, vals2 *[32]uint8, ret *[32]uint8)")
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

	Comment("Packed unsigned byte max per lane (VEX, per 128-bit lane)")
	VPMAXUB(regY2, regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()
	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
