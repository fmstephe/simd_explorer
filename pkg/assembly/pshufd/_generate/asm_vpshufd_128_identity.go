package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufd_128_identity.go -out ../asm_vpshufd_128_identity.s -stubs ../stub_vpshufd_128_identity.go -pkg pshufd
func main() {
	TEXT("vpshufd128Identity", NOSPLIT, "func(vals *[4]uint32, ret *[4]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	reg := XMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFD imm8=0xE4 (identity: [0,1,2,3])")
	VPSHUFD(U8(0xE4), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
