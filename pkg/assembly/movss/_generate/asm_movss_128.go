package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_movss_128.go -out ../asm_movss_128.s -stubs ../stub_movss_128.go -pkg movss
func main() {
	TEXT("movss128", NOSPLIT, "func(vals *[4]float32, ret *[4]float32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX := XMM()
	MOVSS(Mem{Base: vals}, regX)

	Comment("Move vals into another XMM register")
	regX2 := XMM()
	MOVSS(regX, regX2)

	Comment("Write contents of the second XMM register into memory region")
	MOVSS(regX2, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
