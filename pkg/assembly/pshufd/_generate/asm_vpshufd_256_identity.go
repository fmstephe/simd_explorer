package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufd_256_identity.go -out ../asm_vpshufd_256_identity.s -stubs ../stub_vpshufd_256_identity.go -pkg pshufd
func main() {
	TEXT("vpshufd256Identity", NOSPLIT, "func(vals *[8]uint32, ret *[8]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register (per-lane)")
	reg := YMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFD imm8=0xE4 (identity per 128-bit lane)")
	VPSHUFD(U8(0xE4), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
