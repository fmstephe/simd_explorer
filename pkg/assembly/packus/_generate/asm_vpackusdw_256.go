package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpackusdw_256.go -out ../asm_vpackusdw_256.s -stubs ../stub_vpackusdw_256.go -pkg packus
func main() {
	TEXT("vpackusdw256", NOSPLIT, "func(vals1 *[8]int32, vals2 *[8]int32, ret *[16]int16)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into YMM registers")
	reg1 := YMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := YMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Pack signed dwords to unsigned words with saturation (per 128-bit lane)")
	VPACKUSDW(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
