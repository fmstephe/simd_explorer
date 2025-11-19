package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpinsrw_128_seven_idx.go -out ../asm_vpinsrw_128_seven_idx.s -stubs ../stub_vpinsrw_128_seven_idx.go -pkg pinsrw
func main() {
	TEXT("vpinsrw128Seven_idx", NOSPLIT, "func(base *[8]uint16, scalar *uint16, ret *[8]uint16)")
	Comment("load params")
	base := Load(Param("base"), GP64())
	scalar := Load(Param("scalar"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load base into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: base}, regX1)

	Comment("Insert 16-bit word at index 7 (VEX) from memory into XMM")
	VPINSRW(U8(0x07), Mem{Base: scalar}, regX1, regX1)

	Comment("Write result into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
