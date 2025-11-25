package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vbroadcasti128_256.go -out ../asm_vbroadcasti128_256.s -stubs ../stub_vbroadcasti128_256.go -pkg vbroadcasti128
func main() {
	TEXT("vbroadcasti128256", NOSPLIT, "func(vals *[2]uint64, ret *[4]uint64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Broadcast 128-bit valss from memory into both 128-bit lanes of YMM")
	regY := YMM()
	// Use VBROADCASTI128 if available; broadcast raw 128-bit valss to both lanes.
	// If assembler lacks VBROADCASTI128, VBROADCASTF128 has identical bitwise effect.
	VBROADCASTI128(Mem{Base: vals}, regY)

	Comment("Write contents of YMM register into memory region")
	VMOVDQU(regY, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()
	RET()

	// generate!
	Generate()
}
