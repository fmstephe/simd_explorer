package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpalignr_128_sixteen.go -out ../asm_vpalignr_128_sixteen.s -stubs ../stub_vpalignr_128_sixteen.go -pkg palignr
func main() {
	TEXT("vpalignr128Sixteen", NOSPLIT, "func(vals1 *[16]uint8, vals2 *[16]uint8, ret *[16]uint8)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into XMM registers")
	reg1 := XMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := XMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Align right by 16 bytes across vals1/vals2 (imm8=16)")
	VPALIGNR(U8(16), reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
