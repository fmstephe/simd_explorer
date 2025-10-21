package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vmovmskpd_128.go -out ../asm_vmovmskpd_128.s -stubs ../stub_vmovmskpd_128.go -pkg movmskpd
func main() {
	TEXT("vmovmskpd128", NOSPLIT, "func(vals *[2]float64, ret *[4]byte)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	VMOVAPD(Mem{Base: vals}, regX1)

	Comment("Extract sign mask values from vals")
	reg32 := GP32()
	VMOVMSKPD(regX1, reg32)

	Comment("Write sign mask values into return memory address")
	MOVL(reg32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
