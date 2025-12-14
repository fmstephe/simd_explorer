package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmovsxdq_128.go -out ../asm_vpmovsxdq_128.s -stubs ../stub_vpmovsxdq_128.go -pkg vpmovsx
func main() {
	TEXT("vpmovsxdq128", NOSPLIT, "func(vals *[4]int32, ret *[2]int64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	valsX := XMM()
	VMOVDQU(Mem{Base: vals}, valsX)

	retX := XMM()

	Comment("Execute the instruction being demonstrated")
	VPMOVSXDQ(valsX, retX)

	Comment("Write results into return memory address")
	VMOVDQU(retX, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
