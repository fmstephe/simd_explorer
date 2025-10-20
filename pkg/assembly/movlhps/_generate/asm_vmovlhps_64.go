package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmovlhps_64.go -out ../asm_vmovlhps_64.s -stubs ../stub_vmovlhps_64.go -pkg movlhps
func main() {
	TEXT("vmovlhps64", NOSPLIT, "func(vals *[2]float32, ret *[2]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into lower half of XMM register")
	regX1 := XMM()
	MOVLPS(Mem{Base: vals}, regX1)

	Comment("Move lower half of XMM register into upper half of another register")
	regX2 := XMM()
	VMOVLHPS(regX1, regX2, regX2)

	Comment("Write upper half of XMM register into return memory address")
	MOVHPS(regX2, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
