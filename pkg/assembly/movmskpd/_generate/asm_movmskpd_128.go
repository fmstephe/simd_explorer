package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_movmskpd_128.go -out ../asm_movmskpd_128.s -stubs ../stub_movmskpd_128.go -pkg movmskpd
func main() {
	TEXT("movmskpd128", NOSPLIT, "func(vals *[2]float64, ret *[4]byte)")
	Comment("load params")
	vals := Load(Param("vals"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals into XMM register")
	regX1 := XMM()
	MOVAPD(Mem{Base: vals}, regX1)

	Comment("Extract sign mask values from vals")
	reg32 := GP32()
	MOVMSKPD(regX1, reg32)

	Comment("Write sign mask values into return memory address")
	MOVL(reg32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
