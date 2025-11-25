package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpbroadcastq_128.go -out ../asm_vpbroadcastq_128.s -stubs ../stub_vpbroadcastq_128.go -pkg vpbroadcastq
func main() {
	TEXT("vpbroadcastq128", NOSPLIT, "func(q uint64, ret *[2]uint64)")
	// generate!

	Comment("load params")
	q := Load(Param("q"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move q into an XMM register to work with VPBROADCASTQ instruction")
	regXB := XMM()
	MOVQ(q, regXB)

	Comment("Broadcast q into XMM register")
	regX := XMM()
	VPBROADCASTQ(regXB, regX)

	Comment("Write contents of XMM register into memory region")
	VMOVDQU(regX, Mem{Base: ret})

	RET()

	Generate()
}
