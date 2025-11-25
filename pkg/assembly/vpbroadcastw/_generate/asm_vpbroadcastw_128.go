package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpbroadcastw_128.go -out ../asm_vpbroadcastw_128.s -stubs ../stub_vpbroadcastw_128.go -pkg vpbroadcastw
func main() {
	TEXT("vpbroadcastw128", NOSPLIT, "func(w uint16, ret *[8]uint16)")
	// generate!

	Comment("load params")
	w := Load(Param("w"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Need to move w into an XMM register to work with VPBROADCASTW instruction")
	regXB := XMM()
	MOVQ(w, regXB)

	Comment("Broadcast w into XMM register")
	regX := XMM()
	VPBROADCASTW(regXB, regX)

	Comment("Write contents of XMM register into memory region")
	VMOVDQU(regX, Mem{Base: ret})

	RET()

	Generate()
}
