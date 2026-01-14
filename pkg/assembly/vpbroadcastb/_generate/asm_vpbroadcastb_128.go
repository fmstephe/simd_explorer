package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpbroadcastb_128.go -out ../asm_vpbroadcastb_128.s -stubs ../stub_vpbroadcastb_128.go -pkg vpbroadcastb
func main() {
	TEXT("vpbroadcastb128", NOSPLIT, "func(b byte, ret *[16]byte)")
	Comment("load b into a 64 bit register, required for load into XMM register")
	b := Load(Param("b"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load b into XMM register, required for broadcast")
	bX := XMM()
	MOVQ(b, bX)

	retX := XMM()

	Comment("Execute the instruction being demonstrated")
	VPBROADCASTB(bX, retX)

	Comment("Write results into return memory address")
	VMOVDQU(retX, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
