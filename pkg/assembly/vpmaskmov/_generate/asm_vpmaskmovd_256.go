package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaskmovd_256.go -out ../asm_vpmaskmovd_256.s -stubs ../stub_vpmaskmovd_256.go -pkg vpmaskmov
func main() {
	TEXT("vpmaskmovd256", NOSPLIT, "func(vals *[8]uint32, mask *[8]uint32, ret *[8]uint32)")
	Comment("Load parameters")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source (vals) and mask into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Masked store of packed dwords: store elements where mask sign-bit is set")
	VPMASKMOVD(regVals, regMask, Mem{Base: ret})

	Comment("Clear upper halves for AVX->SSE transition")
	VZEROUPPER()

	Comment("Return")
	RET()

	// generate!
	Generate()
}
