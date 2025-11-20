package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vbroadcastss_128_m32.go -out ../asm_vbroadcastss_128_m32.s -stubs ../stub_vbroadcastss_128_m32.go -pkg vbroadcast
func main() {
	TEXT("vbroadcastss128M32", NOSPLIT, "func(scalar *float32, ret *[4]float32)")
	Comment("load params")
	scalar := Load(Param("scalar"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Broadcast 32-bit scalar from memory to all lanes of XMM")
	regX1 := XMM()
	VBROADCASTSS(Mem{Base: scalar}, regX1)

	Comment("Write result into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
