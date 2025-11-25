package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpbroadcastd_256.go -out ../asm_vpbroadcastd_256.s -stubs ../stub_vpbroadcastd_256.go -pkg vpbroadcastd
func main() {
	TEXT("vpbroadcastd256", NOSPLIT, "func(d uint32, ret *[8]uint32)")

	Comment("load params")
	d := Load(Param("d"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move d into an XMM register to work with VPBROADCASTD instruction")
	regXB := XMM()
	MOVQ(d, regXB)

	Comment("Broadcast d into YMM register")
	regY := YMM()
	VPBROADCASTD(regXB, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQU(regY, Mem{Base: ret})

	Comment("Call VZEROUPPER to avoid performance problems after AVX work")
	VZEROUPPER()
	RET()

	// generate!
	Generate()
}
