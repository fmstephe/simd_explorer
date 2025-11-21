package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermilps_128_identity.go -out ../asm_vpermilps_128_identity.s -stubs ../stub_vpermilps_128_identity.go -pkg vpermilps
func main() {
	TEXT("vpermilps128Identity", NOSPLIT, "func(vals *[4]float32, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into XMM")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("VPERMILPS imm8=0xE4 (identity): dst = [a0 a1 a2 a3]")
	VPERMILPS(U8(0xE4), regX1, regX1)

	Comment("Store result")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
