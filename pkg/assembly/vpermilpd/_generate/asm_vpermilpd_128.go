package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermilpd_128.go -out ../asm_vpermilpd_128.s -stubs ../stub_vpermilpd_128.go -pkg vpermilpd
func main() {
	TEXT("vpermilpd128", NOSPLIT, "func(vals *[2]float64, control *[2]float64, ret *[2]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	control := Load(Param("control"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source and control into XMM registers")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals}, regX1)
	regCtrl := XMM()
	VMOVDQA(Mem{Base: control}, regCtrl)

	Comment("VPERMILPD with control register (per-lane 2-bit selectors)")
	VPERMILPD(regCtrl, regX1, regX1)

	Comment("Store result")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
