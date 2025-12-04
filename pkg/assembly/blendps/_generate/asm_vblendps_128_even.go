package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendps_128_even.go -out ../asm_vblendps_128_even.s -stubs ../stub_vblendps_128_even.go -pkg blendps
func main() {
	TEXT("vblendps128Even", NOSPLIT, "func(base *[4]float32, blend *[4]float32, ret *[4]float32)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base and blend into XMM registers")
	regBase := XMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := XMM()
	VMOVDQA(Mem{Base: blend}, regBlend)

	Comment("VBLENDPS 128: imm=0x05 selects even lanes (0,2) from blend; others from base")
	VBLENDPS(U8(0x05), regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
