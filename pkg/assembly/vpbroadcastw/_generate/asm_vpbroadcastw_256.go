package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpbroadcastw_256.go -out ../asm_vpbroadcastw_256.s -stubs ../stub_vpbroadcastw_256.go -pkg vpbroadcastw
func main() {
	TEXT("vpbroadcastw256", NOSPLIT, "func(w uint16, ret *[16]uint16)")

	Comment("load params")
	w := Load(Param("w"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move w into an XMM register to work with VPBROADCASTW instruction")
	regXB := XMM()
	MOVQ(w, regXB)

	Comment("Broadcast w into YMM register")
	regY := YMM()
	VPBROADCASTW(regXB, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQU(regY, Mem{Base: ret})

	Comment("Call VZEROUPPER to avoid performance problems after AVX work")
	VZEROUPPER()
	RET()

	// generate!
	Generate()
}
