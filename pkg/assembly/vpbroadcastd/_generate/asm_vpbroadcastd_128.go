package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpbroadcastd_128.go -out ../asm_vpbroadcastd_128.s -stubs ../stub_vpbroadcastd_128.go -pkg vpbroadcastd
func main() {
	TEXT("vpbroadcastd128", NOSPLIT, "func(d uint32, ret *[4]uint32)")
	// generate!

	Comment("load params")
	d := Load(Param("d"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move d into an XMM register to work with VPBROADCASTD instruction")
	regXB := XMM()
	MOVQ(d, regXB)

	Comment("Broadcast d into XMM register")
	regX := XMM()
	VPBROADCASTD(regXB, regX)

	Comment("Write contents of XMM register into memory region")
	VMOVDQU(regX, Mem{Base: ret})

	RET()

	Generate()
}
