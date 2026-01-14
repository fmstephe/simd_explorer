package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpbroadcastb_256.go -out ../asm_vpbroadcastb_256.s -stubs ../stub_vpbroadcastb_256.go -pkg vpbroadcastb
func main() {
	TEXT("vpbroadcastb256", NOSPLIT, "func(b byte, ret *[32]byte)")

	Comment("load b into 64 bit register, required for load in XMM register")
	b := Load(Param("b"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	bX := XMM()
	MOVQ(b, bX)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VPBROADCASTB(bX, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
