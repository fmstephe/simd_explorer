package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsllw_128.go -out ../asm_vpsllw_128.s -stubs ../stub_vpsllw_128.go -pkg psll
func main() {
	TEXT("vpsllw128", NOSPLIT, "func(vals *[8]uint16, shift *[2]uint64, ret *[8]uint16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shift := Load(Param("shift"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load values and shift counts into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShift := XMM()
	VMOVDQA(Mem{Base: shift}, regShift)

	Comment("Logical left shift packed words by per-lane counts (register)")
	VPSLLW(regShift, regVals, regVals)

	Comment("Write results into return memory address")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
