package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpcmpeqd_128.go -out ../asm_vpcmpeqd_128.s -stubs ../stub_vpcmpeqd_128.go -pkg pcmpeq
func main() {
	TEXT("vpcmpeqd128", NOSPLIT, "func(vals1 *[4]uint32, vals2 *[4]uint32, ret *[4]uint32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load into XMM registers")
	reg1 := XMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := XMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Compare packed dwords for equality; result lanes are 0xFFFFFFFF or 0x00000000")
	VPCMPEQD(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
