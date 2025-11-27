package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpgatherdq_128.go -out ../asm_vpgatherdq_128.s -stubs ../stub_vpgatherdq_128.go -pkg vpgatherdq
func main() {
	TEXT("vpgatherdq128", NOSPLIT, "func(base *[4]uint64, index *[4]uint32, mask *[2]uint64, ret *[2]uint64)")
	Comment("Load parameters")
	base := Load(Param("base"), GP64())
	index := Load(Param("index"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load 32-bit indices into XMM register (lower 2 qwords used)")
	regIndex := XMM()
	VMOVDQA(Mem{Base: index}, regIndex)

	Comment("Load mask into XMM register (MSB of each dword lane used)")
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Load destination/src vector")
	regDst := XMM()
	VMOVDQA(Mem{Base: ret}, regDst)

	Comment("Gather using 32-bit indices; element size is 4 bytes (dword)")
	VPGATHERDQ(regMask, Mem{Base: base, Index: regIndex, Scale: 8}, regDst)

	Comment("Store result vector")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("Return")
	RET()

	// generate!
	Generate()
}
