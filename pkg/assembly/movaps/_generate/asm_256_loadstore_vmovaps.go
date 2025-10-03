package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_256_loadstore_vmovaps.go -out ../asm_256_loadstore_vmovaps.s -stubs ../stub_256_loadstore_vmovaps.go -pkg movaps
func main() {
	TEXT("movaps256LoadStoreVmovaps", NOSPLIT, "func(vals *[8]float32, ret *[8]float32)")

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

	RET()

	// generate!
	Generate()
}
