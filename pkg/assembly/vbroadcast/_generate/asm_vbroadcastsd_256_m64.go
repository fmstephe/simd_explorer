package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vbroadcastsd_256_m64.go -out ../asm_vbroadcastsd_256_m64.s -stubs ../stub_vbroadcastsd_256_m64.go -pkg vbroadcast
func main() {
	TEXT("vbroadcastsd256M64", NOSPLIT, "func(scalar *float64, ret *[4]float64)")
	Comment("load params")
	scalar := Load(Param("scalar"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Broadcast 64-bit scalar from memory to all lanes of YMM (4 lanes)")
	regY1 := YMM()
	VBROADCASTSD(Mem{Base: scalar}, regY1)

	Comment("Write result into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
