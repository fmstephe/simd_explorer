package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpgatherqq_128.go -out ../asm_vpgatherqq_128.s -stubs ../stub_vpgatherqq_128.go -pkg vpgatherqq
func main() {
	TEXT("vpgatherqq128", NOSPLIT, "func(base *[8]uint64, index *[2]uint64, mask *[2]uint64, ret *[2]uint64)")
	Comment("Load parameters")
	base := Load(Param("base"), GP64())
	index := Load(Param("index"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load 64-bit indices into XMM register (lower 2 qwords used)")
	regIndex := XMM()
	VMOVDQA(Mem{Base: index}, regIndex)

	Comment("Load mask into XMM register (MSB of each qword lane used)")
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Load destination/src vector")
	regDst := XMM()
	VMOVDQA(Mem{Base: ret}, regDst)

	Comment("Gather using 64-bit indices; element size is 8 bytes (qword)")
	VPGATHERQQ(regMask, Mem{Base: base, Index: regIndex, Scale: 8}, regDst)

	Comment("Store result vector")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("Return")
	RET()

	// generate!
	Generate()
}
