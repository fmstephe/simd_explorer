package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaskmovd_128.go -out ../asm_vpmaskmovd_128.s -stubs ../stub_vpmaskmovd_128.go -pkg vpmaskmov
func main() {
	TEXT("vpmaskmovd128", NOSPLIT, "func(vals *[4]uint32, mask *[4]uint32, ret *[4]uint32)")
	Comment("Load parameters")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source (vals) and mask into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Masked store of packed dwords: store elements where mask sign-bit is set")
	VPMASKMOVD(regVals, regMask, Mem{Base: ret})

	Comment("Return")
	RET()

	// generate!
	Generate()
}
