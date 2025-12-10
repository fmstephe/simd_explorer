package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vroundpd_256_one.go -out ../asm_vroundpd_256_one.s -stubs ../stub_vroundpd_256_one.go -pkg roundp
func main() {
	TEXT("vroundpd256One", NOSPLIT, "func(vals *[4]float64, ret *[4]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	reg := YMM()
	VMOVAPD(Mem{Base: vals}, reg)

	Comment("Round packed doubles imm8=1 (floor)")
	VROUNDPD(U8(0x01), reg, reg)

	Comment("Write results into return memory address")
	VMOVAPD(reg, Mem{Base: ret})
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
