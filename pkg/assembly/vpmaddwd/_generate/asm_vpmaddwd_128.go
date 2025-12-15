package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaddwd_128.go -out ../asm_vpmaddwd_128.s -stubs ../stub_vpmaddwd_128.go -pkg vpmaddwd
func main() {
	TEXT("vpmaddwd128", NOSPLIT, "func(vals1 *[8]int16, vals2 *[8]int16, ret *[4]int32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into XMM register")
	vals1X := XMM()
	VMOVDQU(Mem{Base: vals1}, vals1X)
	Comment("Load vals2 into XMM register")
	vals2X := XMM()
	VMOVDQU(Mem{Base: vals2}, vals2X)

	retX := XMM()

	Comment("Execute the instruction being demonstrated")
	VPMADDWD(vals1X, vals2X, retX)

	Comment("Write results into return memory address")
	VMOVDQU(retX, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
