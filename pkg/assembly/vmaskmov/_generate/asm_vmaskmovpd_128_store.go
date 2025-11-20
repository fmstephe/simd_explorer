package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmaskmovpd_128_store.go -out ../asm_vmaskmovpd_128_store.s -stubs ../stub_vmaskmovpd_128_store.go -pkg vmaskmov
func main() {
	TEXT("vmaskmovpd128Store", NOSPLIT, "func(vals *[2]float64, mask *[2]float64, ret *[2]float64)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source (vals) and mask into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Masked store: store selected double-precision elements to memory (ret)")
	VMASKMOVPD(regMask, regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
