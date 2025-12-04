package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendvpw_128.go -out ../asm_vblendvpw_128.s -stubs ../stub_vblendvpw_128.go -pkg blendvpw
func main() {
	TEXT("vblendvpw128", NOSPLIT, "func(base *[8]uint16, blend *[8]uint16, mask *[8]uint16, ret *[8]uint16)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base, blend, and mask into XMM registers")
	regBase := XMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := XMM()
	VMOVDQA(Mem{Base: blend}, regBlend)
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("VPBLENDVB: variable blend using per-byte sign bits of mask")
	VPBLENDVB(regMask, regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
