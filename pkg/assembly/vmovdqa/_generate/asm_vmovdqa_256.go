package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmovdqa_256.go -out ../asm_vmovdqa_256.s -stubs ../stub_vmovdqa_256.go -pkg vmovdqa
func main() {
	TEXT("vmovdqa256", NOSPLIT, "func(vals *[8]uint32, ret *[8]uint32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	regY := YMM()
	VMOVDQA(Mem{Base: vals}, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQA(regY, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	RET()

	// generate!
	Generate()
}
