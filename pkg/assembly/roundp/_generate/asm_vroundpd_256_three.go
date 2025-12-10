package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vroundpd_256_three.go -out ../asm_vroundpd_256_three.s -stubs ../stub_vroundpd_256_three.go -pkg roundp
func main() {
	TEXT("vroundpd256Three", NOSPLIT, "func(vals *[4]float64, ret *[4]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	reg := YMM()
	VMOVAPD(Mem{Base: vals}, reg)

	Comment("Round packed doubles imm8=3 (truncate)")
	VROUNDPD(U8(0x03), reg, reg)

	Comment("Write results into return memory address")
	VMOVAPD(reg, Mem{Base: ret})
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
