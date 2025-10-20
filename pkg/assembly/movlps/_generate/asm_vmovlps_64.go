package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmovlps_64.go -out ../asm_vmovlps_64.s -stubs ../stub_vmovlps_64.go -pkg movlps
func main() {
	TEXT("vmovlps64", NOSPLIT, "func(lower, upper *[2]float32, ret *[4]float32)")

	Comment("load params")
	lower := Load(Param("lower"), GP64())
	upper := Load(Param("upper"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load upper (64 bit) into upper half of XMM register")
	regUpper := XMM()
	MOVHPS(Mem{Base: upper}, regUpper)

	Comment("Load lower (64 bit) and merge with upper in register")
	regMerged := XMM()
	VMOVLPS(Mem{Base: lower}, regUpper, regMerged)

	Comment("Write contents of the merged XMM register into memory region")
	VMOVUPS(regMerged, Mem{Base: ret})

	RET()

	// generate!
	Generate()
}
