package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmovsxbd_256.go -out ../asm_vpmovsxbd_256.s -stubs ../stub_vpmovsxbd_256.go -pkg vpmovsx
func main() {
	TEXT("vpmovsxbd256", NOSPLIT, "func(vals *[16]int8, ret *[8]int32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	valsX := XMM()
	VMOVDQU(Mem{Base: vals}, valsX)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VPMOVSXBD(valsX, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
