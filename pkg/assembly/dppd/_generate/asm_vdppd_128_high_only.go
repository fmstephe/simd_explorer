package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vdppd_128_high_only.go -out ../asm_vdppd_128_high_only.s -stubs ../stub_vdppd_128_high_only.go -pkg dppd
func main() {
	TEXT("vdppd128High_only", NOSPLIT, "func(vals1 *[2]float64, vals2 *[2]float64, ret *[2]float64)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into XMM registers")
	reg1 := XMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := XMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Dot product of packed doubles with imm8=0x23 (write high only)")
	VDPPD(U8(0x23), reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
