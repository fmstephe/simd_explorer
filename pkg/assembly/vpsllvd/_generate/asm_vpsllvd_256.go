package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsllvd_256.go -out ../asm_vpsllvd_256.s -stubs ../stub_vpsllvd_256.go -pkg vpsllvd
func main() {
	TEXT("vpsllvd256", NOSPLIT, "func(vals *[8]uint32, shifts *[8]uint32, ret *[8]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shifts := Load(Param("shifts"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals and shifts into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShifts := YMM()
	VMOVDQA(Mem{Base: shifts}, regShifts)

	Comment("VPSLLVD: shift packed doubleword integers in vals left by variable counts in shifts")
	VPSLLVD(regShifts, regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
