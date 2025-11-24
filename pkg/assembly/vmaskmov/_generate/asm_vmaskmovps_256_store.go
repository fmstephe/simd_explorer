package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmaskmovps_256_store.go -out ../asm_vmaskmovps_256_store.s -stubs ../stub_vmaskmovps_256_store.go -pkg vmaskmov
func main() {
	TEXT("vmaskmovps256Store", NOSPLIT, "func(vals *[8]float32, mask *[8]float32, ret *[8]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load source (vals) and mask into YMM registers")
	regVals := YMM()
	VMOVDQA(Mem{Base: vals}, regVals)
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)

	Comment("Masked store: store selected single-precision elements to memory (ret)")
	VMASKMOVPS(regMask, regVals, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
