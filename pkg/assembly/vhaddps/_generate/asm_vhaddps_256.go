package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vhaddps_256.go -out ../asm_vhaddps_256.s -stubs ../stub_vhaddps_256.go -pkg vhaddps
func main() {
	TEXT("vhaddps256", NOSPLIT, "func(vals1 *[8]float32, vals2 *[8]float32, ret *[8]float32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into YMM register")
	vals1Y := YMM()
	VMOVDQU(Mem{Base: vals1}, vals1Y)
	Comment("Load vals2 into YMM register")
	vals2Y := YMM()
	VMOVDQU(Mem{Base: vals2}, vals2Y)

	retY := YMM()

	Comment("Execute the instruction being demonstrated")
	VHADDPS(vals2Y, vals1Y, retY)

	Comment("Write results into return memory address")
	VMOVDQU(retY, Mem{Base: ret})

	Comment("Clear upper halves after YMM usage")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
