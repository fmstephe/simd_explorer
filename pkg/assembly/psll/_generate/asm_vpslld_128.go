package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpslld_128.go -out ../asm_vpslld_128.s -stubs ../stub_vpslld_128.go -pkg psll
func main() {
	TEXT("vpslld128", NOSPLIT, "func(vals *[4]uint32, shift *[2]uint64, ret *[4]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shift := Load(Param("shift"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load values and shift counts into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShift := XMM()
	VMOVDQA(Mem{Base: shift}, regShift)

	Comment("Logical left shift packed dwords by per-lane counts (register)")
	VPSLLD(regShift, regVals, regVals)

	Comment("Write results into return memory address")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
