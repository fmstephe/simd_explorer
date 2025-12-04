package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendps_256_odd.go -out ../asm_vblendps_256_odd.s -stubs ../stub_vblendps_256_odd.go -pkg blendps
func main() {
	TEXT("vblendps256Odd", NOSPLIT, "func(base *[8]float32, blend *[8]float32, ret *[8]float32)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base and blend into YMM registers")
	regBase := YMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := YMM()
	VMOVDQA(Mem{Base: blend}, regBlend)

	Comment("VBLENDPS 256: imm=0xAA selects odd lanes (1,3,5,7) from blend; others from base")
	VBLENDPS(U8(0xAA), regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
