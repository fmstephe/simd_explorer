package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmaskmovps_256_load.go -out ../asm_vmaskmovps_256_load.s -stubs ../stub_vmaskmovps_256_load.go -pkg vmaskmov
func main() {
	TEXT("vmaskmovps256Load", NOSPLIT, "func(vals *[8]float32, mask *[8]float32, ret *[8]float32)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	mask := Load(Param("mask"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load mask into YMM register")
	regMask := YMM()
	VMOVDQA(Mem{Base: mask}, regMask)
	Comment("Zero destination")
	regDst := YMM()
	VXORPS(regDst, regDst, regDst)

	Comment("Masked load: load selected elements from memory into destination")
	VMASKMOVPS(Mem{Base: vals}, regMask, regDst)

	Comment("Write results into return memory address")
	VMOVDQA(regDst, Mem{Base: ret})

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
