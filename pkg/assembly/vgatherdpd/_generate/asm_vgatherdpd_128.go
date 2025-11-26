package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vgatherdpd_128.go -out ../asm_vgatherdpd_128.s -stubs ../stub_vgatherdpd_128.go -pkg vgatherdpd
func main() {
	TEXT("vgatherdpd128", NOSPLIT, "func(base *[8]float64, index *[4]uint32, mask *[2]float64, ret *[2]float64)")
	Comment("Load parameters")
	base := Load(Param("base"), GP64())
	index := Load(Param("index"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load index (i32) into XMM index register")
	regIndex := XMM()
	VMOVDQA(Mem{Base: index}, regIndex)

	Comment("Load mask (packed f64; MSB of each element used)")
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Load destination/src vector (merge for masked-off lanes)")
	regDst := XMM()
	VMOVDQA(Mem{Base: ret}, regDst)

	Comment("Gather two f64 elements using 32-bit indices; scale=8 for f64 element size")
	VGATHERDPD(regMask, Mem{Base: base, Index: regIndex, Scale: 8}, regDst)

	Comment("Store result vector")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("Return")
	RET()

	// generate!
	Generate()
}
