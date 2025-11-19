package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_pmovmskb_128.go -out ../asm_pmovmskb_128.s -stubs ../stub_pmovmskb_128.go -pkg pmovmskb
func main() {
	TEXT("pmovmskb128", NOSPLIT, "func(vals *[16]uint8, ret *uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Move MSBs of bytes to mask (lower 16 bits of 32-bit reg)")
	reg32 := GP32()
	PMOVMSKB(regX1, reg32)

	Comment("Write mask to ret")
	MOVL(reg32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
