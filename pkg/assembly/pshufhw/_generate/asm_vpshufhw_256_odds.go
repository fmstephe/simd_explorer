package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufhw_256_odds.go -out ../asm_vpshufhw_256_odds.s -stubs ../stub_vpshufhw_256_odds.go -pkg pshufhw
func main() {
	TEXT("vpshufhw256Odds", NOSPLIT, "func(vals *[16]uint16, ret *[16]uint16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register (per-lane)")
	reg := YMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFHW imm8=0xDD (odds per lane)")
	VPSHUFHW(U8(0xDD), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
