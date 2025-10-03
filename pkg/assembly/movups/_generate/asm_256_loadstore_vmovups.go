package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_256_loadstore_vmovups.go -out ../asm_256_loadstore_vmovups.s -stubs ../stub_256_loadstore_vmovups.go -pkg movups
func main() {
	TEXT("movups256LoadStoreVmovups", NOSPLIT, "func(vals *[8]float32, ret *[8]float32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regY := YMM()
	VMOVUPS(Mem{Base: vals}, regY)

	Comment("Move vals into another YMM register")
	regY2 := YMM()
	VMOVUPS(regY, regY2)

	Comment("Write contents of the second YMM register into memory region")
	VMOVUPS(regY2, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
