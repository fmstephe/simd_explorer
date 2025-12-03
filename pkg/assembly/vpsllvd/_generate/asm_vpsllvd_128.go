package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsllvd_128.go -out ../asm_vpsllvd_128.s -stubs ../stub_vpsllvd_128.go -pkg vpsllvd
func main() {
	TEXT("vpsllvd128", NOSPLIT, "func(vals *[4]uint32, shifts *[4]uint32, ret *[4]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shifts := Load(Param("shifts"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals and shifts into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShifts := XMM()
	VMOVDQA(Mem{Base: shifts}, regShifts)

	Comment("VPSLLVD: shift packed doubleword integers in vals left by variable counts in shifts")
	VPSLLVD(regShifts, regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
