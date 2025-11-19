package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmovmskb_128.go -out ../asm_vpmovmskb_128.s -stubs ../stub_vpmovmskb_128.go -pkg pmovmskb
func main() {
	TEXT("vpmovmskb128", NOSPLIT, "func(vals *[16]uint8, ret *uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)

	Comment("Move MSBs of bytes to mask (lower 16 bits of 32-bit reg) VEX")
	reg32 := GP32()
	VPMOVMSKB(regX1, reg32)

	Comment("Write mask to ret")
	MOVL(reg32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
