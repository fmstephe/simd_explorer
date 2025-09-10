package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm.go -out ../asm.s -stubs ../stub.go -pkg vpbroadcastb
func main() {
	TEXT("vpbroadcastb", NOSPLIT, "func(b byte, ret *[256]byte)")
	// generate!

	Comment("load params")
	b := Load(Param("b"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move b into an XMM register to work with VPBROADCASTB instruction")
	regX := XMM()
	MOVQ(b, regX)

	Comment("Broadcast b into YMM register")
	regY := YMM()
	VPBROADCASTB(regX, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQU(regY, Mem{Base: ret})

	Comment("Call VZEROUPPER to avoid performance problems after AVX work")
	VZEROUPPER()
	RET()

	Generate()
}
