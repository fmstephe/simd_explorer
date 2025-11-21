package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermilpd_256.go -out ../asm_vpermilpd_256.s -stubs ../stub_vpermilpd_256.go -pkg vpermilpd
func main() {
	TEXT("vpermilpd256", NOSPLIT, "func(vals *[4]float64, control *[4]float64, ret *[4]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	control := Load(Param("control"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source and control into YMM registers")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals}, regY1)
	regCtrl := YMM()
	VMOVDQA(Mem{Base: control}, regCtrl)

	Comment("VPERMILPD with control register (per-lane 2-bit selectors, per 128-bit lane)")
	VPERMILPD(regCtrl, regY1, regY1)

	Comment("Store result")
	VMOVDQA(regY1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
