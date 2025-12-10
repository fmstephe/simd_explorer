package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vroundps_128_three.go -out ../asm_vroundps_128_three.s -stubs ../stub_vroundps_128_three.go -pkg roundp
func main() {
	TEXT("vroundps128Three", NOSPLIT, "func(vals *[4]float32, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	reg := XMM()
	VMOVAPS(Mem{Base: vals}, reg)

	Comment("Round packed singles imm8=3 (truncate)")
	VROUNDPS(U8(0x03), reg, reg)

	Comment("Write results into return memory address")
	VMOVAPS(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
