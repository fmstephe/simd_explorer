package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vblendpd_128_none.go -out ../asm_vblendpd_128_none.s -stubs ../stub_vblendpd_128_none.go -pkg blendpd
func main() {
	TEXT("vblendpd128None", NOSPLIT, "func(base *[2]float64, blend *[2]float64, ret *[2]float64)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base and blend into XMM registers")
	regBase := XMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := XMM()
	VMOVDQA(Mem{Base: blend}, regBlend)

	Comment("VBLENDPD 128: imm=0x00 selects all lanes from base")
	VBLENDPD(U8(0x00), regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
