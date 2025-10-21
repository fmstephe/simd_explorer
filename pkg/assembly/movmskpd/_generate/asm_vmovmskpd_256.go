package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmovmskpd_256.go -out ../asm_vmovmskpd_256.s -stubs ../stub_vmovmskpd_256.go -pkg movmskpd
func main() {
	TEXT("vmovmskpd256", NOSPLIT, "func(vals *[4]float64, ret *[4]byte)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into YMM register")
	regY1 := YMM()
	VMOVAPD(Mem{Base: vals}, regY1)

	Comment("Extract sign mask values from vals")
	reg32 := GP32()
	VMOVMSKPD(regY1, reg32)

	Comment("Write sign mask values into return memory address")
	MOVL(reg32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
