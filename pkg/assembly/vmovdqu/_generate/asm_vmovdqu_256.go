package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmovdqu_256.go -out ../asm_vmovdqu_256.s -stubs ../stub_vmovdqu_256.go -pkg vmovdqu
func main() {
	TEXT("vmovdqu256", NOSPLIT, "func(vals *[8]uint32, ret *[8]uint32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	regY := YMM()
	VMOVDQU(Mem{Base: vals}, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQU(regY, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
