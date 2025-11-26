package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vgatherqpd_256.go -out ../asm_vgatherqpd_256.s -stubs ../stub_vgatherqpd_256.go -pkg vgatherqdp
func main() {
	TEXT("vgatherqpd256", NOSPLIT, "func(base *[8]float64, index *[4]uint64, mask *[4]float64, ret *[4]float64)")
	Comment("Load parameters")
	base := Load(Param("base"), GP64())
	index := Load(Param("index"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load 64-bit indices into YMM register (lower 4 qwords used)")
	regIndex := YMM()
	VMOVDQA(Mem{Base: index}, regIndex)

	Comment("Load mask into YMM register")
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Load destination/src vector")
	regDst := YMM()
	VMOVDQA(Mem{Base: ret}, regDst)

	Comment("Gather four f64 elements using 64-bit indices; scale=8 for f64 element size")
	VGATHERQPD(regMask, Mem{Base: base, Index: regIndex, Scale: 8}, regDst)

	Comment("Store result vector")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return")
	RET()

	// generate!
	Generate()
}
