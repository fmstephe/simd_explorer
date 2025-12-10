package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vroundsd_128_zero.go -out ../asm_vroundsd_128_zero.s -stubs ../stub_vroundsd_128_zero.go -pkg rounds
func main() {
	TEXT("vroundsd128Zero", NOSPLIT, "func(base *[2]float64, vals *[2]float64, ret *[2]float64)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base and vals into XMM registers")
	regBase := XMM()
	VMOVAPD(Mem{Base: base}, regBase)
	regVals := XMM()
	VMOVAPD(Mem{Base: vals}, regVals)

	Comment("Round scalar double imm8=0 (nearest), copy upper lane from base")
	VROUNDSD(U8(0x00), regVals, regBase, regBase)

	Comment("Write results into return memory address")
	VMOVAPD(regBase, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
