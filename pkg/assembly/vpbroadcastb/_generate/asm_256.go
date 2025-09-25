package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_256.go -out ../asm_256.s -stubs ../stub_256.go -pkg vpbroadcastb
func main() {
	TEXT("vpbroadcastb256", NOSPLIT, "func(b byte, ret *[32]byte)")
	// generate!

	Comment("load params")
	b := Load(Param("b"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	regXArg := XMM()
	MOVQ(b, regXArg)

	Comment("Broadcast b into YMM register")
	regY := YMM()
	VPBROADCASTB(regXArg, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQU(regY, Mem{Base: ret})

	Comment("Call VZEROUPPER to avoid performance problems after AVX work")
	VZEROUPPER()
	RET()
	Generate()
}
