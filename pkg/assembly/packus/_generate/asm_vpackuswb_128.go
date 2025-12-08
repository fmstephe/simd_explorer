package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpackuswb_128.go -out ../asm_vpackuswb_128.s -stubs ../stub_vpackuswb_128.go -pkg packus
func main() {
	TEXT("vpackuswb128", NOSPLIT, "func(vals1 *[8]int16, vals2 *[8]int16, ret *[16]int8)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into XMM registers")
	reg1 := XMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := XMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Pack signed words to unsigned bytes with saturation")
	VPACKUSWB(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
