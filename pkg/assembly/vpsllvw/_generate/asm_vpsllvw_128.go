package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsllvw_128.go -out ../asm_vpsllvw_128.s -stubs ../stub_vpsllvw_128.go -pkg vpsllvw
func main() {
	TEXT("vpsllvw128", NOSPLIT, "func(vals *[8]uint16, shifts *[8]uint16, ret *[8]uint16)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shifts := Load(Param("shifts"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals and shifts into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShifts := XMM()
	VMOVDQA(Mem{Base: shifts}, regShifts)

	Comment("VPSLLVW: shift packed doubleword integers in vals left by variable counts in shifts")
	VPSLLVW(regShifts, regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
