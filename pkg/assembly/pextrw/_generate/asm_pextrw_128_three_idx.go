package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_pextrw_128_three_idx.go -out ../asm_pextrw_128_three_idx.s -stubs ../stub_pextrw_128_three_idx.go -pkg pextrw
func main() {
	TEXT("pextrw128Three_idx", NOSPLIT, "func(vals *[8]uint16, ret *uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Extract 16-bit word at index 3 into 32-bit GPR (zero-extended)")
	reg32 := GP32()
	PEXTRW(U8(0x03), regX1, reg32)

	Comment("Write result into return memory address")
	MOVL(reg32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
