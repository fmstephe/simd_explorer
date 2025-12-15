package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpabsw_256.go -out ../asm_vpabsw_256.s -stubs ../stub_vpabsw_256.go -pkg vpabs
func main() {
	TEXT("vpabsw256", NOSPLIT, "func(vals *[16]int16, ret *[16]int16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	valsY := YMM()
	VMOVDQU(Mem{Base: vals}, valsY)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VPABSW(valsY, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
