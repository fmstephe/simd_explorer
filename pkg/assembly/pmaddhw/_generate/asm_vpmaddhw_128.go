package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmaddhw_128.go -out ../asm_vpmaddhw_128.s -stubs ../stub_vpmaddhw_128.go -pkg pmaddhw
func main() {
	TEXT("vpmaddhw128", NOSPLIT, "func(vals1 *[8]int16, vals2 *[8]int16, ret *[4]int32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM registers")
	reg1 := XMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := XMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("VPMADDWD: multiply signed 16-bit pairs and add adjacent products -> 32-bit results")
	VPMADDWD(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
