package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermilps_128_all_three.go -out ../asm_vpermilps_128_all_three.s -stubs ../stub_vpermilps_128_all_three.go -pkg vpermilps
func main() {
	TEXT("vpermilps128All_three", NOSPLIT, "func(vals *[4]float32, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into XMM")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("VPERMILPS imm8=0xFF: broadcast a3 to all lanes")
	VPERMILPS(U8(0xFF), regX1, regX1)

	Comment("Store result")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
