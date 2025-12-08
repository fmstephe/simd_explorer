package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsrlw_256.go -out ../asm_vpsrlw_256.s -stubs ../stub_vpsrlw_256.go -pkg psrl
func main() {
	TEXT("vpsrlw256", NOSPLIT, "func(vals *[16]uint16, shift *[4]uint64, ret *[16]uint16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shift := Load(Param("shift"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load values and shift counts into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShift := XMM()
	VMOVDQA(Mem{Base: shift}, regShift)

	Comment("Logical right shift packed words by per-lane counts (register)")
	VPSRLW(regShift, regVals, regVals)

	Comment("Write results into return memory address")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
