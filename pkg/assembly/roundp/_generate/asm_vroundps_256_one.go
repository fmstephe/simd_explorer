package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vroundps_256_one.go -out ../asm_vroundps_256_one.s -stubs ../stub_vroundps_256_one.go -pkg roundp
func main() {
	TEXT("vroundps256One", NOSPLIT, "func(vals *[8]float32, ret *[8]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	reg := YMM()
	VMOVAPS(Mem{Base: vals}, reg)

	Comment("Round packed singles imm8=1 (floor)")
	VROUNDPS(U8(0x01), reg, reg)

	Comment("Write results into return memory address")
	VMOVAPS(reg, Mem{Base: ret})
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
