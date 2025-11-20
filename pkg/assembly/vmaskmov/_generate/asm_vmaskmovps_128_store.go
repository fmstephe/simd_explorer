package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmaskmovps_128_store.go -out ../asm_vmaskmovps_128_store.s -stubs ../stub_vmaskmovps_128_store.go -pkg vmaskmov
func main() {
	TEXT("vmaskmovps128Store", NOSPLIT, "func(vals *[4]float32, mask *[4]float32, ret *[4]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source (vals) and mask into XMM registers")
	regVals := XMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regMask := XMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Masked store: store selected single-precision elements to memory (ret)")
	VMASKMOVPS(regMask, regVals, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
