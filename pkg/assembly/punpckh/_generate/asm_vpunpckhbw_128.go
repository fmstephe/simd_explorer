package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpunpckhbw_128.go -out ../asm_vpunpckhbw_128.s -stubs ../stub_vpunpckhbw_128.go -pkg punpckh
func main() {
	TEXT("vpunpckhbw128", NOSPLIT, "func(vals1 *[16]uint8, vals2 *[16]uint8, ret *[16]uint8)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into XMM registers")
	reg1 := XMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := XMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Interleave high-order bytes from each operand")
	VPUNPCKHBW(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
