package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpgatherdq_256.go -out ../asm_vpgatherdq_256.s -stubs ../stub_vpgatherdq_256.go -pkg vpgatherdq
func main() {
	TEXT("vpgatherdq256", NOSPLIT, "func(base *[8]uint64, index *[4]uint32, mask *[4]uint64, ret *[4]uint64)")
	Comment("Load parameters")
	base := Load(Param("base"), GP64())
	index := Load(Param("index"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load 32-bit indices into XMM register (lower 4 qwords used)")
	regIndex := XMM()
	VMOVDQA(Mem{Base: index}, regIndex)

	Comment("Load mask into YMM register (MSB of each dword lane used)")
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Load destination/src vector")
	regDst := YMM()
	VMOVDQA(Mem{Base: ret}, regDst)

	Comment("Gather using 32-bit indices; element size is 4 bytes (dword)")
	VPGATHERDQ(regMask, Mem{Base: base, Index: regIndex, Scale: 8}, regDst)

	Comment("Store result vector")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return")
	RET()

	// generate!
	Generate()
}
