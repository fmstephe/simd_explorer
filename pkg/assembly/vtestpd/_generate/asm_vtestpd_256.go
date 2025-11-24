package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vtestpd_256.go -out ../asm_vtestpd_256.s -stubs ../stub_vtestpd_256.go -pkg vtestpd
func main() {
	TEXT("vtestpd256", NOSPLIT, "func(vals1, vals2 *[4]float64, ret *uint32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load inputs into YMM registers")
	regY1 := YMM()
	VMOVDQA(Mem{Base: vals1}, regY1)
	regY2 := YMM()
	VMOVDQA(Mem{Base: vals2}, regY2)

	Comment("VTESTPD sets ZF/CF based on sign-bit AND/ANDN across lanes (per 128-bit lane)")
	VTESTPD(regY1, regY2)

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

	Comment("YMM/ZMM processing complete, clear upper half of YMM registers")
	VZEROUPPER()

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
