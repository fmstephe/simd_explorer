package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufd_128_evens.go -out ../asm_vpshufd_128_evens.s -stubs ../stub_vpshufd_128_evens.go -pkg pshufd
func main() {
	TEXT("vpshufd128Evens", NOSPLIT, "func(vals *[4]uint32, ret *[4]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	reg := XMM()
	VMOVDQA(Mem{Base: vals}, reg)

	Comment("VPSHUFD imm8=0x88 (evens: [0,2,0,2])")
	VPSHUFD(U8(0x88), reg, reg)

	Comment("Write results into return memory address")
	VMOVDQA(reg, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
