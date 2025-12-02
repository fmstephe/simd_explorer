package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermq_256_all_ones.go -out ../asm_vpermq_256_all_ones.s -stubs ../stub_vpermq_256_all_ones.go -pkg vpermq
func main() {
	TEXT("vpermq256All_ones", NOSPLIT, "func(vals *[4]uint64, ret *[4]uint64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into YMM register")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)

	Comment("VPERMQ all ones: imm8=0x55 selects index 1 for all outputs")
	VPERMQ(Imm(0x55), regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
