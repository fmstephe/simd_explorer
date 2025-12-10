package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vroundps_128_two.go -out ../asm_vroundps_128_two.s -stubs ../stub_vroundps_128_two.go -pkg roundp
func main() {
	TEXT("vroundps128Two", NOSPLIT, "func(vals *[4]float32, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	reg := XMM()
	VMOVAPS(Mem{Base: vals}, reg)

	Comment("Round packed singles imm8=2 (ceil)")
	VROUNDPS(U8(0x02), reg, reg)

	Comment("Write results into return memory address")
	VMOVAPS(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
