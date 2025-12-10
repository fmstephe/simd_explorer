package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vroundpd_256_zero.go -out ../asm_vroundpd_256_zero.s -stubs ../stub_vroundpd_256_zero.go -pkg roundp
func main() {
	TEXT("vroundpd256Zero", NOSPLIT, "func(vals *[4]float64, ret *[4]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	reg := YMM()
	VMOVAPD(Mem{Base: vals}, reg)

	Comment("Round packed doubles imm8=0 (nearest)")
	VROUNDPD(U8(0x00), reg, reg)

	Comment("Write results into return memory address")
	VMOVAPD(reg, Mem{Base: ret})
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
