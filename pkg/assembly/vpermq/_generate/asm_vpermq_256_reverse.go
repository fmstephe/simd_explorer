package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermq_256_reverse.go -out ../asm_vpermq_256_reverse.s -stubs ../stub_vpermq_256_reverse.go -pkg vpermq
func main() {
	TEXT("vpermq256Reverse", NOSPLIT, "func(vals *[4]uint64, ret *[4]uint64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into YMM register")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)

	Comment("VPERMQ reverse: imm8=0x1B selects indices 3,2,1,0")
	VPERMQ(Imm(0x1B), regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
