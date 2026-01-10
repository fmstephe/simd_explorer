package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufb_256.go -out ../asm_vpshufb_256.s -stubs ../stub_vpshufb_256.go -pkg pshufb
func main() {
	TEXT("vpshufb256", NOSPLIT, "func(vals1 *[32]uint8, control *[32]uint8, ret *[32]uint8)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	control := Load(Param("control"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into YMM register")
	vals1Y := YMM()
	VMOVDQU(Mem{Base: vals1}, vals1Y)
	Comment("Load control into YMM register")
	controlY := YMM()
	VMOVDQU(Mem{Base: control}, controlY)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VPSHUFB(controlY, vals1Y, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
