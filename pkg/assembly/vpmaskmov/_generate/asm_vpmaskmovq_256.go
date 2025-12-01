package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaskmovq_256.go -out ../asm_vpmaskmovq_256.s -stubs ../stub_vpmaskmovq_256.go -pkg vpmaskmov
func main() {
	TEXT("vpmaskmovq256", NOSPLIT, "func(vals *[4]uint64, mask *[4]uint64, ret *[4]uint64)")
	Comment("Load parameters")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source (vals) and mask into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Masked store of packed qwords: store elements where mask sign-bit is set")
	VPMASKMOVQ(regVals, regMask, Mem{Base: ret})

	Comment("Clear upper halves for AVX->SSE transition")
	VZEROUPPER()

	Comment("Return")
	RET()

	// generate!
	Generate()
}
