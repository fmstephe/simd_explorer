package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshuflw_128_odds.go -out ../asm_vpshuflw_128_odds.s -stubs ../stub_vpshuflw_128_odds.go -pkg pshuflw
func main() {
	TEXT("vpshuflw128Odds", NOSPLIT, "func(vals *[8]uint16, ret *[8]uint16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	reg := XMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFLW imm8=0xDD (odds: [w5,w7,w5,w7])")
	VPSHUFLW(U8(0xDD), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
