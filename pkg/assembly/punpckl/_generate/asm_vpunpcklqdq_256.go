package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpunpcklqdq_256.go -out ../asm_vpunpcklqdq_256.s -stubs ../stub_vpunpcklqdq_256.go -pkg punpckl
func main() {
	TEXT("vpunpcklqdq256", NOSPLIT, "func(vals1 *[4]uint64, vals2 *[4]uint64, ret *[4]uint64)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into YMM registers (per-lane operation)")
	reg1 := YMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := YMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Interleave low-order qwords from each operand (per lane)")
	VPUNPCKLQDQ(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
