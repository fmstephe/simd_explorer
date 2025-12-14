package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmovzxbw_128.go -out ../asm_vpmovzxbw_128.s -stubs ../stub_vpmovzxbw_128.go -pkg vpmovzx
func main() {
	TEXT("vpmovzxbw128", NOSPLIT, "func(vals *[16]uint8,  ret *[8]uint16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	valsX := XMM()
	VMOVDQU(Mem{Base: vals}, valsX)

	retX := XMM()

	Comment("Execute the instruction being demonstrated")
	VPMOVZXBW(valsX, retX)

	Comment("Write results into return memory address")
	VMOVDQU(retX, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
