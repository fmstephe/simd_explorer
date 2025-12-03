package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vpsllvq_256.go -out ../asm_vpsllvq_256.s -stubs ../stub_vpsllvq_256.go -pkg vpsllvq
func main() {
	TEXT("vpsllvq256", NOSPLIT, "func(vals *[4]uint64, shifts *[4]uint64, ret *[4]uint64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	shifts := Load(Param("shifts"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals and shifts into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regShifts := YMM()
	VMOVDQA(Mem{Base: shifts}, regShifts)

	Comment("VPSLLVQ: shift packed quadword integers in vals left by variable counts in shifts")
	VPSLLVQ(regShifts, regVals, regVals)

	Comment("Store result")
	VMOVDQA(regVals, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
