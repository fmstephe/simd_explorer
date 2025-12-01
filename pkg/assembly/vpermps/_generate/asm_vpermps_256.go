package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpermps_256.go -out ../asm_vpermps_256.s -stubs ../stub_vpermps_256.go -pkg vpermps
func main() {
	TEXT("vpermps256", NOSPLIT, "func(vals *[8]float32, control *[8]uint32, ret *[8]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	control := Load(Param("control"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load control and source into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regCtrl := YMM()
	VMOVDQA(Mem{Base: control}, regCtrl)

	Comment("VPERMPS: permute packed single-precision floats with per-lane dword indices")
	VPERMPS(regVals, regCtrl, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
