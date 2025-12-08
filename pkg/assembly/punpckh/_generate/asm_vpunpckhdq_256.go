package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpunpckhdq_256.go -out ../asm_vpunpckhdq_256.s -stubs ../stub_vpunpckhdq_256.go -pkg punpckh
func main() {
	TEXT("vpunpckhdq256", NOSPLIT, "func(vals1 *[8]uint32, vals2 *[8]uint32, ret *[8]uint32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into YMM registers (per-lane operation)")
	reg1 := YMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := YMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Interleave high-order doublewords from each operand (per lane)")
	VPUNPCKHDQ(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
