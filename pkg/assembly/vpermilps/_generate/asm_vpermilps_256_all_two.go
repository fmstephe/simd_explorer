package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermilps_256_all_two.go -out ../asm_vpermilps_256_all_two.s -stubs ../stub_vpermilps_256_all_two.go -pkg vpermilps
func main() {
	TEXT("vpermilps256All_two", NOSPLIT, "func(vals *[8]float32, ret *[8]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into YMM")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals}, regY1)

	Comment("VPERMILPS imm8=0xAA per 128-bit lane: broadcast lane2 element")
	VPERMILPS(U8(0xAA), regY1, regY1)

	Comment("Store result")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
