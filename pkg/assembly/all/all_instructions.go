package all

import (
	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/assembly/addps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/addss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/andnps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/andps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cmpps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cmpss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/comiss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/divps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/divss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/maxps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/maxss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/minps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/minss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movaps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movhps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movlhps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movlps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movmskpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movmskps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/movups"
	"github.com/fmstephe/simd_explorer/pkg/assembly/mulps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/mulss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/orps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pavgb"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pavgw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pextrw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pinsrw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmaxsw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmaxub"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pminsw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pminub"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmovmskb"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmulhuw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/psadbw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rcpps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rcpss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rsqrtps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rsqrtss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/shufps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/sqrtps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/sqrtss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/subps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/subss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/ucomiss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/unpckhps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vbroadcast"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vbroadcasti128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vextractf128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vinsertf128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmaskmov"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmovdqa"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmovdqu"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastb"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vperm2f128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpermilps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vtestpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vtestps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/xorps"
)

func Instructions() []assembly.Instruction {
	return []assembly.Instruction{
		// vpbroadcastb
		&vpbroadcastb.VPBROADCASTB128{},
		&vpbroadcastb.VPBROADCASTB256{},
		&vpbroadcastb.VPBROADCASTB512{},
		&vpbroadcastb.VPBROADCASTB128K{},
		&vpbroadcastb.VPBROADCASTB256K{},
		&vpbroadcastb.VPBROADCASTB512K{},
		// vpbroadcastw
		&vpbroadcastw.VPBROADCASTW128{},
		&vpbroadcastw.VPBROADCASTW256{},
		// vpbroadcastd
		&vpbroadcastd.VPBROADCASTD128{},
		&vpbroadcastd.VPBROADCASTD256{},
		// vpbroadcastq
		&vpbroadcastq.VPBROADCASTQ128{},
		&vpbroadcastq.VPBROADCASTQ256{},
		// vbroadcasti128
		&vbroadcasti128.VBROADCASTI128256{},
		// vmovdqu
		&vmovdqu.VMOVDQU128{},
		&vmovdqu.VMOVDQU256{},
		&vmovdqa.VMOVDQA128{},
		&vmovdqa.VMOVDQA256{},
		// addps
		&addps.ADDPS128{},
		&addps.VADDPS128{},
		&addps.VADDPS256{},
		// mulps
		&mulps.MULPS128{},
		&mulps.VMULPS128{},
		&mulps.VMULPS256{},
		// subps
		&subps.SUBPS128{},
		&subps.VSUBPS128{},
		&subps.VSUBPS256{},
		// andps
		&andps.ANDPS128{},
		&andps.VANDPS128{},
		&andps.VANDPS256{},
		// orps
		&orps.ORPS128{},
		&orps.VORPS128{},
		&orps.VORPS256{},
		// xorps
		&xorps.XORPS128{},
		&xorps.VXORPS128{},
		&xorps.VXORPS256{},
		// andnps
		&andnps.ANDNPS128{},
		&andnps.VANDNPS128{},
		&andnps.VANDNPS256{},
		// pmulhuw
		&pmulhuw.PMULHUW128{},
		&pmulhuw.VPMULHUW128{},
		&pmulhuw.VPMULHUW256{},
		// addss
		&addss.ADDSS128{},
		&addss.VADDSS128{},
		// divss
		&divss.DIVSS128{},
		&divss.VDIVSS128{},
		// divps
		&divps.DIVPS128{},
		&divps.VDIVPS128{},
		&divps.VDIVPS256{},
		// subss
		&subss.SUBSS128{},
		&subss.VSUBSS128{},
		// cmpps (128-bit)
		&cmpps.CMPPS128EQ{},
		&cmpps.CMPPS128LT{},
		&cmpps.CMPPS128LE{},
		&cmpps.CMPPS128UNORD{},
		&cmpps.CMPPS128NEQ{},
		&cmpps.CMPPS128NLT{},
		&cmpps.CMPPS128NLE{},
		&cmpps.CMPPS128ORD{},
		// mulss
		&mulss.MULSS128{},
		&mulss.VMULSS128{},
		// movss
		&movss.MOVSS128{},
		&movss.VMOVSS128{},
		// cmpss
		&cmpss.CMPSS128EQ{},
		&cmpss.CMPSS128LT{},
		&cmpss.CMPSS128LE{},
		&cmpss.CMPSS128UNORD{},
		&cmpss.CMPSS128NEQ{},
		&cmpss.CMPSS128NLT{},
		&cmpss.CMPSS128NLE{},
		&cmpss.CMPSS128ORD{},
		&cmpss.VCMPSS128EQ{},
		&cmpss.VCMPSS128LT{},
		&cmpss.VCMPSS128LE{},
		&cmpss.VCMPSS128UNORD{},
		&cmpss.VCMPSS128NEQ{},
		&cmpss.VCMPSS128NLT{},
		&cmpss.VCMPSS128NLE{},
		&cmpss.VCMPSS128ORD{},
		// comiss
		&comiss.COMISS128{},
		&comiss.VCOMISS128{},
		// ucomiss
		&ucomiss.UCOMISS128{},
		&ucomiss.VUCOMISS128{},
		// rcpps
		&rcpps.RCPPS128{},
		&rcpps.VRCPPS128{},
		&rcpps.VRCPPS256{},
		// sqrtps
		&sqrtps.SQRTPS128{},
		&sqrtps.VSQRTPS128{},
		&sqrtps.VSQRTPS256{},
		// rsqrtps
		&rsqrtps.RSQRTPS128{},
		&rsqrtps.VRSQRTPS128{},
		&rsqrtps.VRSQRTPS256{},
		// rcpss
		&rcpss.RCPSS128{},
		&rcpss.VRCPSS128{},
		// rsqrtss
		&rsqrtss.RSQRTSS128{},
		&rsqrtss.VRSQRTSS128{},
		// sqrtss
		&sqrtss.SQRTSS128{},
		&sqrtss.VSQRTSS128{},
		// movaps
		&movaps.MOVAPS128{},
		&movaps.VMOVAPS128{},
		&movaps.VMOVAPS256{},
		// movups
		&movups.MOVUPS128{},
		&movups.VMOVUPS128{},
		&movups.VMOVUPS256{},
		// movlps
		&movlps.MOVLPS64{},
		&movlps.VMOVLPS64{},
		// movhpos
		&movhps.MOVHPS64{},
		&movhps.VMOVHPS64{},
		// unpckhps
		&unpckhps.UNPCKHPS128{},
		&unpckhps.VUNPCKHPS128{},
		&unpckhps.VUNPCKHPS256{},
		// shufps
		&shufps.SHUFPS128ZEROS{},
		&shufps.SHUFPS128ONES{},
		&shufps.SHUFPS128TWOS{},
		&shufps.SHUFPS128THREES{},
		&shufps.SHUFPS128MIXED{},
		&shufps.SHUFPS128REVERSE{},
		&shufps.VSHUFPS128ZEROS{},
		&shufps.VSHUFPS128ONES{},
		&shufps.VSHUFPS128TWOS{},
		&shufps.VSHUFPS128THREES{},
		&shufps.VSHUFPS128MIXED{},
		&shufps.VSHUFPS128REVERSE{},
		&shufps.VSHUFPS256ZEROS{},
		&shufps.VSHUFPS256ONES{},
		&shufps.VSHUFPS256TWOS{},
		&shufps.VSHUFPS256THREES{},
		&shufps.VSHUFPS256MIXED{},
		&shufps.VSHUFPS256REVERSE{},
		// movlhps
		&movlhps.MOVLHPS64{},
		&movlhps.VMOVLHPS64{},
		// maxps
		&maxps.MAXPS128{},
		&maxps.VMAXPS128{},
		&maxps.VMAXPS256{},
		// maxss
		&maxss.MAXSS128{},
		&maxss.VMAXSS128{},
		// minps
		&minps.MINPS128{},
		&minps.VMINPS128{},
		&minps.VMINPS256{},
		// minss
		&minss.MINSS128{},
		&minss.VMINSS128{},
		// movmskps
		&movmskps.MOVMSKPS128{},
		&movmskps.VMOVMSKPS128{},
		&movmskps.VMOVMSKPS256{},
		// movmskpd
		&movmskpd.MOVMSKPD128{},
		&movmskpd.VMOVMSKPD128{},
		&movmskpd.VMOVMSKPD256{},
		// psadbw
		&psadbw.PSADBW128{},
		&psadbw.VPSADBW128{},
		&psadbw.VPSADBW256{},
		// pavgb
		&pavgb.PAVGB128{},
		&pavgb.VPAVGB128{},
		&pavgb.VPAVGB256{},
		// pavgw
		&pavgw.PAVGW128{},
		&pavgw.VPAVGW128{},
		&pavgw.VPAVGW256{},
		// pmaxub
		&pmaxub.PMAXUB128{},
		&pmaxub.VPMAXUB128{},
		&pmaxub.VPMAXUB256{},
		// pminub
		&pminub.PMINUB128{},
		&pminub.VPMINUB128{},
		&pminub.VPMINUB256{},
		// pmaxsw
		&pmaxsw.PMAXSW128{},
		&pmaxsw.VPMAXSW128{},
		&pmaxsw.VPMAXSW256{},
		// pminsw
		&pminsw.PMINSW128{},
		&pminsw.VPMINSW128{},
		&pminsw.VPMINSW256{},
		// pextrw
		&pextrw.PEXTRW128ZERO_IDX{},
		&pextrw.PEXTRW128ONE_IDX{},
		&pextrw.PEXTRW128TWO_IDX{},
		&pextrw.PEXTRW128THREE_IDX{},
		&pextrw.PEXTRW128FOUR_IDX{},
		&pextrw.PEXTRW128FIVE_IDX{},
		&pextrw.PEXTRW128SIX_IDX{},
		&pextrw.PEXTRW128SEVEN_IDX{},
		&pextrw.VPEXTRW128ZERO_IDX{},
		&pextrw.VPEXTRW128ONE_IDX{},
		&pextrw.VPEXTRW128TWO_IDX{},
		&pextrw.VPEXTRW128THREE_IDX{},
		&pextrw.VPEXTRW128FOUR_IDX{},
		&pextrw.VPEXTRW128FIVE_IDX{},
		&pextrw.VPEXTRW128SIX_IDX{},
		&pextrw.VPEXTRW128SEVEN_IDX{},
		// pinsrw
		&pinsrw.PINSRW128ZERO_IDX{},
		&pinsrw.PINSRW128ONE_IDX{},
		&pinsrw.PINSRW128TWO_IDX{},
		&pinsrw.PINSRW128THREE_IDX{},
		&pinsrw.PINSRW128FOUR_IDX{},
		&pinsrw.PINSRW128FIVE_IDX{},
		&pinsrw.PINSRW128SIX_IDX{},
		&pinsrw.PINSRW128SEVEN_IDX{},
		&pinsrw.VPINSRW128ZERO_IDX{},
		&pinsrw.VPINSRW128ONE_IDX{},
		&pinsrw.VPINSRW128TWO_IDX{},
		&pinsrw.VPINSRW128THREE_IDX{},
		&pinsrw.VPINSRW128FOUR_IDX{},
		&pinsrw.VPINSRW128FIVE_IDX{},
		&pinsrw.VPINSRW128SIX_IDX{},
		&pinsrw.VPINSRW128SEVEN_IDX{},
		// pmovmskb
		&pmovmskb.PMOVMSKB128{},
		&pmovmskb.VPMOVMSKB128{},
		// vbroadcast
		&vbroadcast.VBROADCASTSS128M32{},
		&vbroadcast.VBROADCASTSS256M32{},
		&vbroadcast.VBROADCASTSD256M64{},
		&vbroadcast.VBROADCASTSF128256M128{},
		// vinsertf128
		&vinsertf128.VINSERTF128256ZERO{},
		&vinsertf128.VINSERTF128256ONE{},
		// vextractf128
		&vextractf128.VEXTRACTF128256ZERO{},
		&vextractf128.VEXTRACTF128256ONE{},
		// vmaskmov
		&vmaskmov.VMASKMOVPS128LOAD{},
		&vmaskmov.VMASKMOVPD128LOAD{},
		&vmaskmov.VMASKMOVPD256LOAD{},
		&vmaskmov.VMASKMOVPD256LOAD{},
		&vmaskmov.VMASKMOVPS128STORE{},
		&vmaskmov.VMASKMOVPD128STORE{},
		&vmaskmov.VMASKMOVPD256STORE{},
		&vmaskmov.VMASKMOVPD256STORE{},
		// vpermilps
		&vpermilps.VPERMILPS128{},
		&vpermilps.VPERMILPS128IDENTITY{},
		&vpermilps.VPERMILPS128ALL_ZERO{},
		&vpermilps.VPERMILPS128ALL_ONE{},
		&vpermilps.VPERMILPS128ALL_TWO{},
		&vpermilps.VPERMILPS128ALL_THREE{},
		&vpermilps.VPERMILPS128REVERSE{},
		&vpermilps.VPERMILPS256{},
		&vpermilps.VPERMILPS256IDENTITY{},
		&vpermilps.VPERMILPS256ALL_ZERO{},
		&vpermilps.VPERMILPS256ALL_ONE{},
		&vpermilps.VPERMILPS256ALL_TWO{},
		&vpermilps.VPERMILPS256ALL_THREE{},
		&vpermilps.VPERMILPS256REVERSE{},
		// vperm2f128
		&vperm2f128.VPERM2F128128LOWA_HIGHA{},
		&vperm2f128.VPERM2F128128LOWB_HIGHB{},
		&vperm2f128.VPERM2F128128LOWA_HIGHB{},
		&vperm2f128.VPERM2F128128LOWB_HIGHA{},
		&vperm2f128.VPERM2F128128LOWA_LOWA{},
		&vperm2f128.VPERM2F128128HIGHA_HIGHA{},
		&vperm2f128.VPERM2F128128LOWB_LOWB{},
		&vperm2f128.VPERM2F128128HIGHB_HIGHB{},
		&vperm2f128.VPERM2F128128ZEROED_HIGHB{},
		&vperm2f128.VPERM2F128128LOWA_ZEROED{},
		&vperm2f128.VPERM2F128128ZEROED_ZEROED{},
		// vtestps
		&vtestps.VTESTPS128{},
		&vtestps.VTESTPS256{},
		// vtestpd
		&vtestpd.VTESTPD128{},
		&vtestpd.VTESTPD256{},
	}
}

func SupportedInstructions() []assembly.Instruction {
	instructions := Instructions()
	supported := make([]assembly.Instruction, 0, len(instructions))
	for _, inst := range instructions {
		if inst.Supported() {
			supported = append(supported, inst)
		}
	}
	return supported
}
