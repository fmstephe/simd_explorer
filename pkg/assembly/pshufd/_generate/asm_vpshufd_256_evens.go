package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufd_256_evens.go -out ../asm_vpshufd_256_evens.s -stubs ../stub_vpshufd_256_evens.go -pkg pshufd
func main() {
	TEXT("vpshufd256Evens", NOSPLIT, "func(vals *[8]uint32, ret *[8]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register (per-lane)")
	reg := YMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFD imm8=0x88 (evens per 128-bit lane: [0,2,0,2])")
	VPSHUFD(U8(0x88), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
