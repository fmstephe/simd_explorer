package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_64_loadstore_movlps.go -out ../asm_64_loadstore_movlps.s -stubs ../stub_64_loadstore_movlps.go -pkg movlps
func main() {
	TEXT("movlps64LoadStoreMovlps", NOSPLIT, "func(vals *[2]float32, ret *[2]float32)")

	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX := XMM()
	MOVLPS(Mem{Base: vals}, regX)

	Comment("Write contents of the XMM register into memory region")
	MOVLPS(regX, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
