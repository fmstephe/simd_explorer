package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsrlvd_128.go -out ../asm_vpsrlvd_128.s -stubs ../stub_vpsrlvd_128.go -pkg vpsrlvd
func main() {
	TEXT("vpsrlvd128", NOSPLIT, "func(vals *[4]uint32, shifts *[4]uint32, ret *[4]uint32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shifts := Load(Param("shifts"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals and shifts into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShifts := XMM()
	VMOVDQA(Mem{Base: shifts}, regShifts)

	Comment("VPSRLVD: shift packed doubleword integers in vals right by variable counts in shifts")
	VPSRLVD(regShifts, regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
