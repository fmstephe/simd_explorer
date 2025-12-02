package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vperm2i128_256_lowb_higha.go -out ../asm_vperm2i128_256_lowb_higha.s -stubs ../stub_vperm2i128_256_lowb_higha.go -pkg vperm2i128
func main() {
	TEXT("vperm2i128256Lowb_higha", NOSPLIT, "func(valsA *[4]uint64, valsB *[4]uint64, ret *[4]uint64)")
	Comment("load params")
	a := Load(Param("valsA"), GP64())
	b := Load(Param("valsB"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load A and B into YMM registers")
	regA := YMM()
	VMOVDQA(Mem{Base: a}, regA)
	regB := YMM()
	VMOVDQA(Mem{Base: b}, regB)

	Comment("VPERM2I128 imm: low=b.low (10), high=a.high (01) → imm=0x12")
	VPERM2I128(U8(0x12), regB, regA, regA)

	Comment("Store result")
	VMOVDQA(regA, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
