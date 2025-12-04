package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendvpb_256.go -out ../asm_vblendvpb_256.s -stubs ../stub_vblendvpb_256.go -pkg blendvpb
func main() {
	TEXT("vblendvpb256", NOSPLIT, "func(base *[32]uint8, blend *[32]uint8, mask *[32]uint8, ret *[32]uint8)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base, blend, and mask into YMM registers")
	regBase := YMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := YMM()
	VMOVDQA(Mem{Base: blend}, regBlend)
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("VPBLENDVB: variable blend using per-byte sign bits of mask")
	VPBLENDVB(regMask, regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
