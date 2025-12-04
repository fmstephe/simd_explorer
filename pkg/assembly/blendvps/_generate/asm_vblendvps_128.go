package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendvps_128.go -out ../asm_vblendvps_128.s -stubs ../stub_vblendvps_128.go -pkg blendvps
func main() {
	TEXT("vblendvps128", NOSPLIT, "func(base *[4]float32, blend *[4]float32, mask *[4]uint32, ret *[4]float32)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base, blend and mask into XMM registers")
	regBase := XMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := XMM()
	VMOVDQA(Mem{Base: blend}, regBlend)
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("VBLENDVPS 128: blend using per-lane sign-bit mask")
	VBLENDVPS(regMask, regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
