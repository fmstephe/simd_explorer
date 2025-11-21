package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vtestps_128.go -out ../asm_vtestps_128.s -stubs ../stub_vtestps_128.go -pkg vtestps
func main() {
	TEXT("vtestps128", NOSPLIT, "func(vals1, vals2 *[4]float32, ret *uint32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load inputs into XMM registers")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals1}, regX1)
	regX2 := XMM()
	VMOVDQA(Mem{Base: vals2}, regX2)

	Comment("VTESTPS sets ZF/CF based on sign-bit AND/ANDN across lanes")
	VTESTPS(regX1, regX2)

	Comment("Capture ZF (SETEQ) and CF (SETCS) into a 32-bit mask: bit0=ZF, bit1=CF")
	zf := GP8()
	cf := GP8()
	SETEQ(zf)
	SETCS(cf)
	z32 := GP32()
	c32 := GP32()
	MOVBLZX(zf, z32)
	MOVBLZX(cf, c32)
	SHLL(U8(1), c32)
	ORL(z32, c32)
	MOVL(c32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
