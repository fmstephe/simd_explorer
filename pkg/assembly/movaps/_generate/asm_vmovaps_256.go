package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmovaps_256.go -out ../asm_vmovaps_256.s -stubs ../stub_vmovaps_256.go -pkg movaps
func main() {
	TEXT("vmovaps256", NOSPLIT, "func(vals *[8]float32, ret *[8]float32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regY := YMM()
	VMOVAPS(Mem{Base: vals}, regY)

	Comment("Move vals into another YMM register")
	regY2 := YMM()
	VMOVAPS(regY, regY2)

	Comment("Write contents of the second YMM register into memory region")
	VMOVAPS(regY2, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	RET()

	// generate!
	Generate()
}
