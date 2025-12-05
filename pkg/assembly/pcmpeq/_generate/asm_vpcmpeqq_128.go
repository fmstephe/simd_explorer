package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpcmpeqq_128.go -out ../asm_vpcmpeqq_128.s -stubs ../stub_vpcmpeqq_128.go -pkg pcmpeq
func main() {
	TEXT("vpcmpeqq128", NOSPLIT, "func(vals1 *[2]uint64, vals2 *[2]uint64, ret *[2]uint64)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load into XMM registers")
	reg1 := XMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := XMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Compare packed qwords for equality; result lanes are all-ones or zero")
	VPCMPEQQ(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
