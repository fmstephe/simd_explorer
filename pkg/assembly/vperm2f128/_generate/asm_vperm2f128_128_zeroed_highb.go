package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vperm2f128_128_zeroed_highb.go -out ../asm_vperm2f128_128_zeroed_highb.s -stubs ../stub_vperm2f128_128_zeroed_highb.go -pkg vperm2f128
func main() {
	TEXT("vperm2f128128Zeroed_highb", NOSPLIT, "func(valsA, valsB, ret *[8]float32)")
	Comment("load params")
	a := Load(Param("valsA"), GP64())
	b := Load(Param("valsB"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load A and B into YMM registers")
	regA := YMM()
	VMOVDQA(Mem{Base: a}, regA)
	regB := YMM()
	VMOVDQA(Mem{Base: b}, regB)

	Comment("VPERM2F128 imm: low=zero (0100), high=b.high (0011) → imm=0x38")
	VPERM2F128(U8(0x38), regB, regA, regA)

	Comment("Store result")
	VMOVDQA(regA, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
