package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendvpd_128.go -out ../asm_vblendvpd_128.s -stubs ../stub_vblendvpd_128.go -pkg blendvpd
func main() {
	TEXT("vblendvpd128", NOSPLIT, "func(base *[2]float64, blend *[2]float64, mask *[2]uint64, ret *[2]float64)")
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

	Comment("VBLENDVPD 128: blend using per-lane sign-bit mask")
	VBLENDVPD(regMask, regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
