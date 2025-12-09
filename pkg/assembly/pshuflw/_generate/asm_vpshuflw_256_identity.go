package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshuflw_256_identity.go -out ../asm_vpshuflw_256_identity.s -stubs ../stub_vpshuflw_256_identity.go -pkg pshuflw
func main() {
	TEXT("vpshuflw256Identity", NOSPLIT, "func(vals *[16]uint16, ret *[16]uint16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register (per-lane)")
	reg := YMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFLW imm8=0xE4 (identity for low words, per 128-bit lane)")
	VPSHUFLW(U8(0xE4), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
