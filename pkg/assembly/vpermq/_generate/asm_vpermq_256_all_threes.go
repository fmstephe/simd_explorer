package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermq_256_all_threes.go -out ../asm_vpermq_256_all_threes.s -stubs ../stub_vpermq_256_all_threes.go -pkg vpermq
func main() {
	TEXT("vpermq256All_threes", NOSPLIT, "func(vals *[4]uint64, ret *[4]uint64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into YMM register")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)

	Comment("VPERMQ all threes: imm8=0xFF selects index 3 for all outputs")
	VPERMQ(Imm(0xFF), regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
