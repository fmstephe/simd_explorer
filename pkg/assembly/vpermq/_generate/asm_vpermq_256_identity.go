package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermq_256_identity.go -out ../asm_vpermq_256_identity.s -stubs ../stub_vpermq_256_identity.go -pkg vpermq
func main() {
	TEXT("vpermq256Identity", NOSPLIT, "func(vals *[4]uint64, ret *[4]uint64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source into YMM register")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)

	Comment("VPERMQ identity: imm8=0xE4 selects indices 0,1,2,3")
	VPERMQ(Imm(0xE4), regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
