package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmovdqu_128.go -out ../asm_vmovdqu_128.s -stubs ../stub_vmovdqu_128.go -pkg vmovdqu
func main() {
	TEXT("vmovdqu128", NOSPLIT, "func(vals *[4]uint32, ret *[4]uint32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX := XMM()
	VMOVDQU(Mem{Base: vals}, regX)

	Comment("Write contents of XMM register into memory region")
	VMOVDQU(regX, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
