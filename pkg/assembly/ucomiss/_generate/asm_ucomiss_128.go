package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm_ucomiss_128.go -out ../asm_ucomiss_128.s -stubs ../stub_ucomiss_128.go -pkg ucomiss
func main() {
	TEXT("ucomiss128", NOSPLIT, "func(vals1, vals2 *[4]float32) (flags uint32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())

	Comment("Load vals1 into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals1}, regX1)

	Comment("Load vals2 into XMM register")
	regX2 := XMM()
	VMOVDQA(Mem{Base: vals2}, regX2)

	Comment("Compare lowest 32 bit floats from vals1 and vals2, sets EFLAGS as a side effect")
	UCOMISS(regX2, regX1)

	flag := GP8()
	temp := GP32()
	flags := GP32()

	Comment("Set byte if ZF flag is true (set to 1 if compared values are equal)")
	SETEQ(flag)
	MOVBLZX(flag, flags)

	Comment("Set byte if PF flag is true (set to 1 if either values were NaN)")
	SETPS(flag)
	MOVBLZX(flag, temp)
	SHLL(Imm(8), flags)
	ORL(temp, flags)

	Comment("Set byte if CF flag is true (set to 1 if vals1 is less than vals2)")
	SETCS(flag)
	MOVBLZX(flag, temp)
	SHLL(Imm(8), flags)
	ORL(temp, flags)

	Store(flags, ReturnIndex(0))

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
