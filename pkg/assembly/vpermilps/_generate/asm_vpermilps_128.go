package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermilps_128.go -out ../asm_vpermilps_128.s -stubs ../stub_vpermilps_128.go -pkg vpermilps
func main() {
	TEXT("vpermilps128", NOSPLIT, "func(vals *[4]float32, control *[4]float32, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	control := Load(Param("control"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source and control into XMM registers")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)
	regCtrl := XMM()
	VMOVDQA(Mem{Base: control}, regCtrl)

	Comment("VPERMILPS with control register (per-lane 2-bit selectors)")
	VPERMILPS(regCtrl, regX1, regX1)

	Comment("Store result")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
