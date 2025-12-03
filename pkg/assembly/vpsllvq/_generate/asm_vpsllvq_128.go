package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsllvq_128.go -out ../asm_vpsllvq_128.s -stubs ../stub_vpsllvq_128.go -pkg vpsllvq
func main() {
	TEXT("vpsllvq128", NOSPLIT, "func(vals *[2]uint64, shifts *[2]uint64, ret *[2]uint64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shifts := Load(Param("shifts"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals and shifts into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShifts := XMM()
	VMOVDQA(Mem{Base: shifts}, regShifts)

	Comment("VPSLLVQ: shift packed quadword integers in vals left by variable counts in shifts")
	VPSLLVQ(regShifts, regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
