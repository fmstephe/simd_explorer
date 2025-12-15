package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsrldq_128_one.go -out ../asm_vpsrldq_128_one.s -stubs ../stub_vpsrldq_128_one.go -pkg vpsrldq
func main() {
	TEXT("vpsrldq128One", NOSPLIT, "func(vals *[16]uint8,  ret *[16]uint8)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	valsX := XMM()
	VMOVDQU(Mem{Base: vals}, valsX)

	retX := XMM()

	Comment("Execute the instruction being demonstrated")
	VPSRLDQ(Imm(1), valsX, retX)

	Comment("Write results into return memory address")
	VMOVDQU(retX, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
