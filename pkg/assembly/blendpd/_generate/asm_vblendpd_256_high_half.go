package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendpd_256_high_half.go -out ../asm_vblendpd_256_high_half.s -stubs ../stub_vblendpd_256_high_half.go -pkg blendpd
func main() {
	TEXT("vblendpd256High_half", NOSPLIT, "func(base *[4]float64, blend *[4]float64, ret *[4]float64)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base and blend into YMM registers")
	regBase := YMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := YMM()
	VMOVDQA(Mem{Base: blend}, regBlend)

	Comment("VBLENDPD 256: imm=0x0C selects high half (lanes 2,3) from blend; others from base")
	VBLENDPD(U8(0x0C), regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
