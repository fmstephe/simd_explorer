package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpgatherqd_256.go -out ../asm_vpgatherqd_256.s -stubs ../stub_vpgatherqd_256.go -pkg vpgatherqd
func main() {
	TEXT("vpgatherqd256", NOSPLIT, "func(base *[16]uint32, index *[4]uint64, mask *[4]uint32, ret *[4]uint32)")
	Comment("Load parameters")
	base := Load(Param("base"), GP64())
	index := Load(Param("index"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load 64-bit indices into YMM register (lower 4 qwords used)")
	regIndex := YMM()
	VMOVDQA(Mem{Base: index}, regIndex)

	Comment("Load mask into XMM register (MSB of each dword lane used)")
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Load destination/src vector")
	regDst := XMM()
	VMOVDQA(Mem{Base: ret}, regDst)

	Comment("Gather four u32 elements using 64-bit indices; scale=4 for dword element size")
	VPGATHERQD(regMask, Mem{Base: base, Index: regIndex, Scale: 4}, regDst)

	Comment("Store result vector")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return")
	RET()

	// generate!
	Generate()
}
