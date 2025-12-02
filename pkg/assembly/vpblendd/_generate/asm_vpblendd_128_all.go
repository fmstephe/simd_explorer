package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpblendd_128_all.go -out ../asm_vpblendd_128_all.s -stubs ../stub_vpblendd_128_all.go -pkg vpblendd
func main() {
	TEXT("vpblendd128All", NOSPLIT, "func(base *[4]uint32, blend *[4]uint32, ret *[4]uint32)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	blend := Load(Param("blend"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base and blend into XMM registers")
	regBase := XMM()
	VMOVDQA(Mem{Base: base}, regBase)
	regBlend := XMM()
	VMOVDQA(Mem{Base: blend}, regBlend)

	Comment("VPBLENDD: all from blend (imm=0x0F selects all lanes from blend)")
	VPBLENDD(U8(0x0F), regBlend, regBase, regBase)

	Comment("Store result")
	VMOVDQA(regBase, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
