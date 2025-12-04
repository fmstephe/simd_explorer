package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendvpd_256.go -out ../asm_vblendvpd_256.s -stubs ../stub_vblendvpd_256.go -pkg blendvpd
func main() {
	TEXT("vblendvpd256", NOSPLIT, "func(base *[4]float64, blend *[4]float64, mask *[4]uint64, ret *[4]float64)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base, blend and mask into YMM registers")
	regBase := YMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := YMM()
	VMOVDQA(Mem{Base: blend}, regBlend)
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("VBLENDVPD 256: blend using per-lane sign-bit mask")
	VBLENDVPD(regMask, regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
