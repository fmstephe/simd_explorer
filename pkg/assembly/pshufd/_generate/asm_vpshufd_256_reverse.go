package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufd_256_reverse.go -out ../asm_vpshufd_256_reverse.s -stubs ../stub_vpshufd_256_reverse.go -pkg pshufd
func main() {
	TEXT("vpshufd256Reverse", NOSPLIT, "func(vals *[8]uint32, ret *[8]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register (per-lane)")
	reg := YMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFD imm8=0x1B (reverse per 128-bit lane)")
	VPSHUFD(U8(0x1B), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
