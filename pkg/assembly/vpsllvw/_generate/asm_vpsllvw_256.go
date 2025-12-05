package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsllvw_256.go -out ../asm_vpsllvw_256.s -stubs ../stub_vpsllvw_256.go -pkg vpsllvw
func main() {
	TEXT("vpsllvw256", NOSPLIT, "func(vals *[16]uint16, shifts *[16]uint16, ret *[16]uint16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shifts := Load(Param("shifts"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals and shifts into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShifts := YMM()
	VMOVDQA(Mem{Base: shifts}, regShifts)

	Comment("VPSLLVW: shift packed doubleword integers in vals left by variable counts in shifts")
	VPSLLVW(regShifts, regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
