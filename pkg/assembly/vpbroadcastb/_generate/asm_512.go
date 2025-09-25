package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_512.go -out ../asm_512.s -stubs ../stub_512.go -pkg vpbroadcastb
func main() {
	TEXT("vpbroadcastb512", NOSPLIT, "func(b byte, ret *[64]byte)")
	// generate!

	Comment("load params")
	b := Load(Param("b"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	regXArg := XMM()
	MOVQ(b, regXArg)

	Comment("Broadcast b into ZMM register")
	regZ := YMM()
	VPBROADCASTB(regXArg, regZ)

	Comment("Write contents of ZMM register into memory region")
	VMOVDQU64(regZ, Mem{Base: ret})

	Comment("Call VZEROUPPER to avoid performance problems after AVX work")
	VZEROUPPER()
	RET()

	Generate()
}
