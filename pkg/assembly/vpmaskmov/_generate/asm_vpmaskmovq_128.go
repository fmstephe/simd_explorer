package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaskmovq_128.go -out ../asm_vpmaskmovq_128.s -stubs ../stub_vpmaskmovq_128.go -pkg vpmaskmov
func main() {
	TEXT("vpmaskmovq128", NOSPLIT, "func(vals *[2]uint64, mask *[2]uint64, ret *[2]uint64)")
	Comment("Load parameters")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source (vals) and mask into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Masked store of packed qwords: store elements where mask sign-bit is set")
	VPMASKMOVQ(regVals, regMask, Mem{Base: ret})

	Comment("Return")
	RET()

	// generate!
	Generate()
}
