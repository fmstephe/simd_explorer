package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpgatherdd_256.go -out ../asm_vpgatherdd_256.s -stubs ../stub_vpgatherdd_256.go -pkg vpgatherdd
func main() {
	TEXT("vpgatherdd256", NOSPLIT, "func(base *[16]uint32, index *[8]uint32, mask *[8]uint32, ret *[8]uint32)")
	Comment("Load parameters")
	base := Load(Param("base"), GP64())
	index := Load(Param("index"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load 32-bit indices into YMM register (lower 4 dwords; will be used with YMM gather semantics via VSIB)")
	regIndex := YMM()
	VMOVDQA(Mem{Base: index}, regIndex)

	Comment("Load mask into YMM register")
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Load destination/src vector")
	regDst := YMM()
	VMOVDQA(Mem{Base: ret}, regDst)

	Comment("Gather eight u32 elements using 32-bit indices; scale=4 for dword element size")
	VPGATHERDD(regMask, Mem{Base: base, Index: regIndex, Scale: 4}, regDst)

	Comment("Store result vector")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return")
	RET()

	// generate!
	Generate()
}
