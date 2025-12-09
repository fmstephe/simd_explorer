package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpalignr_256_three.go -out ../asm_vpalignr_256_three.s -stubs ../stub_vpalignr_256_three.go -pkg palignr
func main() {
	TEXT("vpalignr256Three", NOSPLIT, "func(vals1 *[32]uint8, vals2 *[32]uint8, ret *[32]uint8)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into YMM registers (per-lane)")
	reg1 := YMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := YMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Align right by 3 bytes across vals1/vals2 per 128-bit lane (imm8=3)")
	VPALIGNR(U8(3), reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
