package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vroundss_128_three.go -out ../asm_vroundss_128_three.s -stubs ../stub_vroundss_128_three.go -pkg rounds
func main() {
	TEXT("vroundss128Three", NOSPLIT, "func(base *[4]float32, vals *[4]float32, ret *[4]float32)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base and vals into XMM registers")
	regBase := XMM()
	VMOVAPS(Mem{Base: base}, regBase)
	regVals := XMM()
	VMOVAPS(Mem{Base: vals}, regVals)

	Comment("Round scalar single imm8=3 (truncate), copy upper lanes from base")
	VROUNDSS(U8(0x03), regVals, regBase, regBase)

	Comment("Write results into return memory address")
	VMOVAPS(regBase, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
