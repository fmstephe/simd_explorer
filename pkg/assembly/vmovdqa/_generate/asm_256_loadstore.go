package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_256_loadstore.go -out ../asm_256_loadstore.s -stubs ../stub_256_loadstore.go -pkg vmovdqu
func main() {
	TEXT("vmovdqu256LoadStore", NOSPLIT, "func(vals *[8]uint32, ret *[8]uint32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	regY := YMM()
	VMOVDQA(Mem{Base: vals}, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQA(regY, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
