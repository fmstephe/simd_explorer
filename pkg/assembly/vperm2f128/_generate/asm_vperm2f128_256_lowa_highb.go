package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vperm2f128_256_lowa_highb.go -out ../asm_vperm2f128_256_lowa_highb.s -stubs ../stub_vperm2f128_256_lowa_highb.go -pkg vperm2f128
func main() {
	TEXT("vperm2f128256Lowa_highb", NOSPLIT, "func(valsA, valsB, ret *[8]float32)")
	Comment("load params")
	a := Load(Param("valsA"), GP64())
	b := Load(Param("valsB"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load A and B into YMM registers")
	regA := YMM()
	VMOVDQA(Mem{Base: a}, regA)
	regB := YMM()
	VMOVDQA(Mem{Base: b}, regB)

	Comment("VPERM2F128 imm: low=a.low (00), high=b.high (11) → imm=0x30")
	VPERM2F128(U8(0x30), regB, regA, regA)

	Comment("Store result")
	VMOVDQA(regA, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
