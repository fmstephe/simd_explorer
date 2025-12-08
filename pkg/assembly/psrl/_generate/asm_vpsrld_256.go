package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsrld_256.go -out ../asm_vpsrld_256.s -stubs ../stub_vpsrld_256.go -pkg psrl
func main() {
	TEXT("vpsrld256", NOSPLIT, "func(vals *[8]uint32, shift *[4]uint64, ret *[8]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shift := Load(Param("shift"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load values and shift counts into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShift := XMM()
	VMOVDQA(Mem{Base: shift}, regShift)

	Comment("Logical right shift packed dwords by per-lane counts (register)")
	VPSRLD(regShift, regVals, regVals)

	Comment("Write results into return memory address")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
