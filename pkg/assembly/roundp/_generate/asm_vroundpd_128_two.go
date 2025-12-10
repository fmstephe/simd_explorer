package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vroundpd_128_two.go -out ../asm_vroundpd_128_two.s -stubs ../stub_vroundpd_128_two.go -pkg roundp
func main() {
	TEXT("vroundpd128Two", NOSPLIT, "func(vals *[2]float64, ret *[2]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	reg := XMM()
	VMOVAPD(Mem{Base: vals}, reg)

	Comment("Round packed doubles imm8=2 (ceil)")
	VROUNDPD(U8(0x02), reg, reg)

	Comment("Write results into return memory address")
	VMOVAPD(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
