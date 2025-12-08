package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpunpckldq_128.go -out ../asm_vpunpckldq_128.s -stubs ../stub_vpunpckldq_128.go -pkg punpckl
func main() {
	TEXT("vpunpckldq128", NOSPLIT, "func(vals1 *[4]uint32, vals2 *[4]uint32, ret *[4]uint32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into XMM registers")
	reg1 := XMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := XMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Interleave low-order dwords from each operand")
	VPUNPCKLDQ(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
