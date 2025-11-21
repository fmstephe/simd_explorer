package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermilpd_128_reverse.go -out ../asm_vpermilpd_128_reverse.s -stubs ../stub_vpermilpd_128_reverse.go -pkg vpermilpd
func main() {
	TEXT("vpermilpd128Reverse", NOSPLIT, "func(vals *[2]float64, ret *[2]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into XMM")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("VPERMILPD imm8=0x1B: reverse [a3 a2 a1 a0]")
	VPERMILPD(U8(0x1B), regX1, regX1)

	Comment("Store result")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
