package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufb_128.go -out ../asm_vpshufb_128.s -stubs ../stub_vpshufb_128.go -pkg pshufb
func main() {
	TEXT("vpshufb128", NOSPLIT, "func(vals1 *[16]uint8, control *[16]uint8, ret *[16]uint8)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	control := Load(Param("control"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into XMM register")
	vals1X := XMM()
	VMOVDQU(Mem{Base: vals1}, vals1X)
	Comment("Load control into XMM register")
	controlX := XMM()
	VMOVDQU(Mem{Base: control}, controlX)

	retX := XMM()

	Comment("Execute the instruction being demonstrated")
	VPSHUFB(controlX, vals1X, retX)

	Comment("Write results into return memory address")
	VMOVDQU(retX, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
