package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermilps_256_all_three.go -out ../asm_vpermilps_256_all_three.s -stubs ../stub_vpermilps_256_all_three.go -pkg vpermilps
func main() {
	TEXT("vpermilps256All_three", NOSPLIT, "func(vals *[8]float32, ret *[8]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into YMM")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals}, regY1)

	Comment("VPERMILPS imm8=0xFF per 128-bit lane: broadcast lane3 element")
	VPERMILPS(U8(0xFF), regY1, regY1)

	Comment("Store result")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
