package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_128_k.go -out ../asm_128_k.s -stubs ../stub_128_k.go -pkg vpbroadcastb
func main() {
	TEXT("vpbroadcastb128K", NOSPLIT, "func(b byte, k uint64, ret *[16]byte)")

	Comment("load params")
	b := Load(Param("b"), GP64())
	k := Load(Param("k"), K())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	regXArg := XMM()
	MOVQ(b, regXArg)

	Comment("Broadcast b into XMM register")
	regX := XMM()
	VPBROADCASTB(regXArg, k, regX)

	Comment("Write contents of XMM register into memory region")
	VMOVDQU(regX, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
