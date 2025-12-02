package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpblendd_128_low_half.go -out ../asm_vpblendd_128_low_half.s -stubs ../stub_vpblendd_128_low_half.go -pkg vpblendd
func main() {
	TEXT("vpblendd128Low_half", NOSPLIT, "func(base *[4]uint32, blend *[4]uint32, ret *[4]uint32)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base and blend into XMM registers")
	regBase := XMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := XMM()
	VMOVDQA(Mem{Base: blend}, regBlend)

	Comment("VPBLENDD: low half from blend (imm=0x03 selects lanes 0 and 1 from blend)")
	VPBLENDD(U8(0x03), regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
