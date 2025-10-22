package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_maxss_128.go -out ../asm_maxss_128.s -stubs ../stub_maxss_128.go -pkg maxss
func main() {
	TEXT("maxss128", NOSPLIT, "func(vals1, vals2 *[4]float32, ret *[4]float32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals1}, regX1)

	Comment("Load vals2 into XMM register")
	regX2 := XMM()
	VMOVDQA(Mem{Base: vals2}, regX2)

	Comment("Compare and return maximum scalar single precision values (lowest 32 bits)")
	MAXSS(regX1, regX2) // Go assembler: regX2 = max(regX2, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX2, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
