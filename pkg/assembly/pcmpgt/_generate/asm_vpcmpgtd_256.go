package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpcmpgtd_256.go -out ../asm_vpcmpgtd_256.s -stubs ../stub_vpcmpgtd_256.go -pkg pcmpgt
func main() {
	TEXT("vpcmpgtd256", NOSPLIT, "func(vals1 *[8]uint32, vals2 *[8]uint32, ret *[8]uint32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load into YMM registers")
	reg1 := YMM()
	VMOVDQA(Mem{Base: vals1}, reg1)
	reg2 := YMM()
	VMOVDQA(Mem{Base: vals2}, reg2)

	Comment("Compare packed signed dwords greater-than (per 128-bit lane)")
	VPCMPGTD(reg2, reg1, reg1)

	Comment("Write results into return memory address")
	VMOVDQA(reg1, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
