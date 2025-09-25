package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_128.go -out ../asm_128.s -stubs ../stub_128.go -pkg vpbroadcastb
func main() {
	TEXT("vpbroadcastb128", NOSPLIT, "func(b byte, ret *[16]byte)")
	// generate!

	Comment("load params")
	b := Load(Param("b"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	regXArg := XMM()
	MOVQ(b, regXArg)

	Comment("Broadcast b into XMM register")
	regX := XMM()
	VPBROADCASTB(regXArg, regX)

	Comment("Write contents of XMM register into memory region")
	VMOVDQU(regX, Mem{Base: ret})

	RET()

	Generate()
}
