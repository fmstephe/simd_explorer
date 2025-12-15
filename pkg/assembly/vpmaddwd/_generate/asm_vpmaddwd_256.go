package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaddwd_256.go -out ../asm_vpmaddwd_256.s -stubs ../stub_vpmaddwd_256.go -pkg vpmaddwd
func main() {
	TEXT("vpmaddwd256", NOSPLIT, "func(vals1 *[16]int16, vals2 *[16]int16, ret *[8]int32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into YMM register")
	vals1Y := YMM()
	VMOVDQU(Mem{Base: vals1}, vals1Y)
	Comment("Load vals2 into YMM register")
	vals2Y := YMM()
	VMOVDQU(Mem{Base: vals2}, vals2Y)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VPMADDWD(vals1Y, vals2Y, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
