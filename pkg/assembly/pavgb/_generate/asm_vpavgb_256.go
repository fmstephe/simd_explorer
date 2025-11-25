package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpavgb_256.go -out ../asm_vpavgb_256.s -stubs ../stub_vpavgb_256.go -pkg pavgb
func main() {
	TEXT("vpavgb256", NOSPLIT, "func(vals1, vals2 *[32]uint8, ret *[32]uint8)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into YMM register")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals1}, regY1)

	Comment("Load vals2 into YMM register")
	regY2 := YMM()
	VMOVDQA(Mem{Base: vals2}, regY2)

	Comment("Average packed unsigned bytes with VEX (per 128-bit lane)")
	VPAVGB(regY2, regY1, regY1)

	Comment("Write results into return memory address")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()
	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
