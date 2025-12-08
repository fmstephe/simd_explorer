package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufd_128_odds.go -out ../asm_vpshufd_128_odds.s -stubs ../stub_vpshufd_128_odds.go -pkg pshufd
func main() {
	TEXT("vpshufd128Odds", NOSPLIT, "func(vals *[4]uint32, ret *[4]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	reg := XMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFD imm8=0xDD (odds: [1,3,1,3])")
	VPSHUFD(U8(0xDD), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
