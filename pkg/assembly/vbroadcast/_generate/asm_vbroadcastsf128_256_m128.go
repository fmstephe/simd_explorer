package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vbroadcastsf128_256_m128.go -out ../asm_vbroadcastsf128_256_m128.s -stubs ../stub_vbroadcastsf128_256_m128.go -pkg vbroadcast
func main() {
	TEXT("vbroadcastsf128256M128", NOSPLIT, "func(block *[4]float32, ret *[8]float32)")
	Comment("load params")
	block := Load(Param("block"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Broadcast 128-bit block from memory to both 128-bit lanes of YMM")
	regY1 := YMM()
	VBROADCASTF128(Mem{Base: block}, regY1)

	Comment("Write result into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
