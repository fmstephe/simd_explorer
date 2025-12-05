package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaddhw_256.go -out ../asm_vpmaddhw_256.s -stubs ../stub_vpmaddhw_256.go -pkg pmaddhw
func main() {
	TEXT("vpmaddhw256", NOSPLIT, "func(vals1 *[16]int16, vals2 *[16]int16, ret *[8]int32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM registers")
	reg1 := YMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := YMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("VPMADDWD: multiply signed 16-bit pairs and add adjacent products -> 32-bit results (per 128-bit lane)")
	VPMADDWD(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
