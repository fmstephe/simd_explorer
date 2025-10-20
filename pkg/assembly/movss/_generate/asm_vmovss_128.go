package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmovss_128.go -out ../asm_vmovss_128.s -stubs ../stub_vmovss_128.go -pkg movss
func main() {
	TEXT("vmovss128", NOSPLIT, "func(vals *[4]float32, ret *[4]float32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX := XMM()
	VMOVSS(Mem{Base: vals}, regX)

	Comment("Move vals into another XMM register")
	regX2 := XMM()
	VMOVSS(regX, regX2, regX2)

	Comment("Write contents of the second XMM register into memory region")
	VMOVSS(regX2, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
