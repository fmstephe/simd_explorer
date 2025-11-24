package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vbroadcastss_256_m32.go -out ../asm_vbroadcastss_256_m32.s -stubs ../stub_vbroadcastss_256_m32.go -pkg vbroadcast
func main() {
	TEXT("vbroadcastss256M32", NOSPLIT, "func(scalar *float32, ret *[8]float32)")
	Comment("load params")
	scalar := Load(Param("scalar"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Broadcast 32-bit scalar from memory to all lanes of YMM")
	regY1 := YMM()
	VBROADCASTSS(Mem{Base: scalar}, regY1)

	Comment("Write result into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
