package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpbroadcastq_256.go -out ../asm_vpbroadcastq_256.s -stubs ../stub_vpbroadcastq_256.go -pkg vpbroadcastq
func main() {
	TEXT("vpbroadcastq256", NOSPLIT, "func(q uint64, ret *[4]uint64)")

	Comment("load params")
	q := Load(Param("q"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move q into an XMM register to work with VPBROADCASTQ instruction")
	regXB := XMM()
	MOVQ(q, regXB)

	Comment("Broadcast q into YMM register")
	regY := YMM()
	VPBROADCASTQ(regXB, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQU(regY, Mem{Base: ret})

	Comment("Call VZEROUPPER to avoid performance problems after AVX work")
	VZEROUPPER()
	RET()

	// generate!
	Generate()
}
