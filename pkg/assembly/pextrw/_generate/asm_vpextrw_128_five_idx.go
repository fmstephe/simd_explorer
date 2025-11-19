package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpextrw_128_five_idx.go -out ../asm_vpextrw_128_five_idx.s -stubs ../stub_vpextrw_128_five_idx.go -pkg pextrw
func main() {
	TEXT("vpextrw128Five_idx", NOSPLIT, "func(vals *[8]uint16, ret *uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Extract 16-bit word at index 5 into 32-bit GPR (zero-extended, VEX)")
	reg32 := GP32()
	VPEXTRW(U8(0x05), regX1, reg32)

	Comment("Write result into return memory address")
	MOVL(reg32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
