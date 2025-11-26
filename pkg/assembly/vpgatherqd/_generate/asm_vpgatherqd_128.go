package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpgatherqd_128.go -out ../asm_vpgatherqd_128.s -stubs ../stub_vpgatherqd_128.go -pkg vpgatherqd
func main() {
	TEXT("vpgatherqd128", NOSPLIT, "func(base *[8]uint32, index *[2]uint64, mask *[4]uint32, ret *[4]uint32)")
	Comment("Load parameters")
	base := Load(Param("base"), GP64())
	index := Load(Param("index"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load 64-bit indices into XMM register (lower 2 qwords used)")
	regIndex := XMM()
	VMOVDQA(Mem{Base: index}, regIndex)

	Comment("Load mask into XMM register (MSB of each dword lane used)")
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Load destination/src vector")
	regDst := XMM()
	VMOVDQA(Mem{Base: ret}, regDst)

	Comment("Gather two u32 elements using 64-bit indices; scale=4 for dword element size")
	VPGATHERQD(regMask, Mem{Base: base, Index: regIndex, Scale: 4}, regDst)

	Comment("Store result vector")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("Return")
	RET()

	// generate!
	Generate()
}
