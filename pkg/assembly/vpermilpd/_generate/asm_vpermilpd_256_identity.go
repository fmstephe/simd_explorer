package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermilpd_256_identity.go -out ../asm_vpermilpd_256_identity.s -stubs ../stub_vpermilpd_256_identity.go -pkg vpermilpd
func main() {
	TEXT("vpermilpd256Identity", NOSPLIT, "func(vals *[4]float64, ret *[4]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into YMM")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals}, regY1)

	Comment("VPERMILPD imm8=0xE4 per 128-bit lane: identity mapping")
	VPERMILPD(U8(0xE4), regY1, regY1)

	Comment("Store result")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
