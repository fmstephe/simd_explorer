package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmovzxdq_128.go -out ../asm_vpmovzxdq_128.s -stubs ../stub_vpmovzxdq_128.go -pkg vpmovzx
func main() {
	TEXT("vpmovzxdq128", NOSPLIT, "func(vals *[4]uint32,  ret *[2]uint64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	valsX := XMM()
	VMOVDQU(Mem{Base: vals}, valsX)

	retX := XMM()

	Comment("Execute the instruction being demonstrated")
	VPMOVZXDQ(valsX, retX)

	Comment("Write results into return memory address")
	VMOVDQU(retX, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
