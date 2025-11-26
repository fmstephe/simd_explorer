package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vgatherqpd_128.go -out ../asm_vgatherqpd_128.s -stubs ../stub_vgatherqpd_128.go -pkg vgatherqdp
func main() {
	TEXT("vgatherqpd128", NOSPLIT, "func(base *[8]float64, index *[2]uint64, mask *[2]float64, ret *[2]float64)")
	Comment("Load parameters")
	base := Load(Param("base"), GP64())
	index := Load(Param("index"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load 64-bit indices into XMM register (lower 2 qwords used)")
	regIndex := XMM()
	VMOVDQA(Mem{Base: index}, regIndex)

	Comment("Load mask into XMM register")
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Load destination/src vector")
	regDst := XMM()
	VMOVDQA(Mem{Base: ret}, regDst)

	Comment("Gather two f64 elements using 64-bit indices; scale=8 for f64 element size")
	VGATHERQPD(regMask, Mem{Base: base, Index: regIndex, Scale: 8}, regDst)

	Comment("Store result vector")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("Return")
	RET()

	// generate!
	Generate()
}
