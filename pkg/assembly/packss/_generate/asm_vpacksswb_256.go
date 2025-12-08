package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpacksswb_256.go -out ../asm_vpacksswb_256.s -stubs ../stub_vpacksswb_256.go -pkg packss
func main() {
	TEXT("vpacksswb256", NOSPLIT, "func(vals1 *[16]int16, vals2 *[16]int16, ret *[32]int8)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into YMM registers")
	reg1 := YMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := YMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Pack signed words to signed bytes with saturation (per 128-bit lane)")
	VPACKSSWB(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
