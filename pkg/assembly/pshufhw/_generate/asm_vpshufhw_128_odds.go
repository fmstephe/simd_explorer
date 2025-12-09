package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufhw_128_odds.go -out ../asm_vpshufhw_128_odds.s -stubs ../stub_vpshufhw_128_odds.go -pkg pshufhw
func main() {
	TEXT("vpshufhw128Odds", NOSPLIT, "func(vals *[8]uint16, ret *[8]uint16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	reg := XMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFHW imm8=0xDD (odds: [w5,w7,w5,w7])")
	VPSHUFHW(U8(0xDD), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
