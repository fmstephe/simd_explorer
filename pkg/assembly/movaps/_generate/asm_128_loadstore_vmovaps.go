package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_128_loadstore_vmovaps.go -out ../asm_128_loadstore_vmovaps.s -stubs ../stub_128_loadstore_vmovaps.go -pkg movaps
func main() {
	TEXT("movaps128LoadStoreVmovaps", NOSPLIT, "func(vals *[4]float32, ret *[4]float32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX := XMM()
	VMOVAPS(Mem{Base: vals}, regX)

	Comment("Move vals into another XMM register")
	regX2 := XMM()
	VMOVAPS(regX, regX2)

	Comment("Write contents of the second XMM register into memory region")
	VMOVAPS(regX2, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
