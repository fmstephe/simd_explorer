package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpshufb_256.go -out ../asm_vpshufb_256.s -stubs ../stub_vpshufb_256.go -pkg pshufb
func main() {
	TEXT("vpshufb256", NOSPLIT, "func(vals1 *[32]uint8, control *[32]uint8, ret *[32]uint8)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	control := Load(Param("control"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load operands into YMM registers (per-lane operation)")
	regData := YMM()
	VMOVDQA(Mem{Base: vals1}, regData)
	regControl := YMM()
	VMOVDQA(Mem{Base: control}, regControl)

	Comment("Shuffle bytes in regData according to regControl (per lane)")
	VPSHUFB(regControl, regData, regData)

	Comment("Write results into return memory address")
	VMOVDQA(regData, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
