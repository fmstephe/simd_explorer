package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vinsertf128_256_one.go -out ../asm_vinsertf128_256_one.s -stubs ../stub_vinsertf128_256_one.go -pkg vinsertf128
func main() {
	TEXT("vinsertf128256One", NOSPLIT, "func(base *[8]float32, block *[4]float32, ret *[8]float32)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	block := Load(Param("block"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base YMM register")
	regY1 := YMM()
	VMOVDQA(Mem{Base: base}, regY1)

	Comment("Insert 128-bit block from memory into lane 1 of YMM (imm8=1)")
	VINSERTF128(U8(0x01), Mem{Base: block}, regY1, regY1)

	Comment("Write result into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
