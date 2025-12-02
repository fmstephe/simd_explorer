package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermq_256_all_zeros.go -out ../asm_vpermq_256_all_zeros.s -stubs ../stub_vpermq_256_all_zeros.go -pkg vpermq
func main() {
	TEXT("vpermq256All_zeros", NOSPLIT, "func(vals *[4]uint64, ret *[4]uint64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into YMM register")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)

	Comment("VPERMQ all zeros: imm8=0x00 selects index 0 for all outputs")
	VPERMQ(Imm(0x00), regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
