package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpmuludq_128.go -out ../asm_vpmuludq_128.s -stubs ../stub_vpmuludq_128.go -pkg pmuludq
func main() {
	TEXT("vpmuludq128", NOSPLIT, "func(vals1 *[4]uint32, vals2 *[4]uint32, ret *[2]uint64)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 and vals2 into XMM registers")
	reg1 := XMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := XMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("VPMULUDQ: multiply pairs of unsigned doublewords, producing 64-bit results in even lanes")
	VPMULUDQ(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
