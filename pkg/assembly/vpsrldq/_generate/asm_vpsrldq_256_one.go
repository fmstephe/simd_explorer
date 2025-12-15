package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsrldq_256_one.go -out ../asm_vpsrldq_256_one.s -stubs ../stub_vpsrldq_256_one.go -pkg vpsrldq
func main() {
	TEXT("vpsrldq256One", NOSPLIT, "func(vals *[32]uint8,  ret *[32]uint8)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	valsY := YMM()
	VMOVDQU(Mem{Base: vals}, valsY)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VPSRLDQ(Imm(1), valsY, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
