package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendvps_256.go -out ../asm_vblendvps_256.s -stubs ../stub_vblendvps_256.go -pkg blendvps
func main() {
	TEXT("vblendvps256", NOSPLIT, "func(base *[8]float32, blend *[8]float32, mask *[8]uint32, ret *[8]float32)")
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

	Comment("VBLENDVPS 256: blend using per-lane sign-bit mask")
	VBLENDVPS(regMask, regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
