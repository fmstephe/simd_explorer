package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsraq_128.go -out ../asm_vpsraq_128.s -stubs ../stub_vpsraq_128.go -pkg psra
func main() {
	TEXT("vpsraq128", NOSPLIT, "func(vals *[2]int64, shift *[2]uint64, ret *[2]int64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shift := Load(Param("shift"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load values and shift counts into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShift := XMM()
	VMOVDQA(Mem{Base: shift}, regShift)

	Comment("Arithmetical right shift packed qwords by per-lane counts (register)")
	VPSRAQ(regShift, regVals, regVals)

	Comment("Write results into return memory address")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
