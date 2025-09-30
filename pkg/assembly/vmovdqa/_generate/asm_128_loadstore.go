package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_128_loadstore.go -out ../asm_128_loadstore.s -stubs ../stub_128_loadstore.go -pkg vmovdqu
func main() {
	TEXT("vmovdqu128LoadStore", NOSPLIT, "func(vals *[4]uint32, ret *[4]uint32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX := XMM()
	VMOVDQA(Mem{Base: vals}, regX)

	Comment("Write contents of XMM register into memory region")
	VMOVDQA(regX, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
