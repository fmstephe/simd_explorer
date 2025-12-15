package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_vptest_128.go -out ../asm_vptest_128.s -stubs ../stub_vptest_128.go -pkg vptest
func main() {
	TEXT("vptest128", NOSPLIT, "func(vals1 *[16]uint8, vals2 *[16]uint8, ret *uint32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into XMM register")
	vals1X := XMM()
	VMOVDQU(Mem{Base: vals1}, vals1X)
	Comment("Load vals2 into XMM register")
	vals2X := XMM()
	VMOVDQU(Mem{Base: vals2}, vals2X)

	Comment("Execute the instruction being demonstrated")
	VPTEST(vals1X, vals2X)

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

	Comment("Store flags into *ret (bit0=ZF, bit1=CF)")
	MOVL(c32, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
