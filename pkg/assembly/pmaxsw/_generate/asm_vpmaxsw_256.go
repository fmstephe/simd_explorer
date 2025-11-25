package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaxsw_256.go -out ../asm_vpmaxsw_256.s -stubs ../stub_vpmaxsw_256.go -pkg pmaxsw
func main() {
	TEXT("vpmaxsw256", NOSPLIT, "func(vals1, vals2 *[16]int16, ret *[16]int16)")
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

	Comment("Packed signed word max per lane (VEX, per 128-bit lane)")
	VPMAXSW(regY2, regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()
	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
