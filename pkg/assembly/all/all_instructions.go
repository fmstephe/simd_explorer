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
	"github.com/fmstephe/simd_explorer/pkg/assembly/vextracti128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vgatherdpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vgatherqdp"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vinsertf128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vinserti128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmaskmov"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmovdqa"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vmovdqu"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpblendd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastb"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vperm2f128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vperm2i128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpermd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpermilpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpermilps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpermps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpermq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpgatherdd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpgatherdq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpgatherqd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpgatherqq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpmaskmov"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpsllvd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpsllvq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpsrlvd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vtestpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vtestps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/xorps"
)

func Instructions() []assembly.Instruction {
	return []assembly.Instruction{
		// vpbroadcastb
		vpbroadcastb.NewVPBROADCASTB128(),
		vpbroadcastb.NewVPBROADCASTB256(),
		vpbroadcastb.NewVPBROADCASTB512(),
		vpbroadcastb.NewVPBROADCASTB128K(),
		vpbroadcastb.NewVPBROADCASTB256K(),
		vpbroadcastb.NewVPBROADCASTB512K(),
		// vpbroadcastw
		vpbroadcastw.NewVPBROADCASTW128(),
		vpbroadcastw.NewVPBROADCASTW256(),
		// vpbroadcastd
		vpbroadcastd.NewVPBROADCASTD128(),
		vpbroadcastd.NewVPBROADCASTD256(),
		// vpbroadcastq
		vpbroadcastq.NewVPBROADCASTQ128(),
		vpbroadcastq.NewVPBROADCASTQ256(),
		// vbroadcasti128
		vbroadcasti128.NewVBROADCASTI128256(),
		// vinserti128
		&vinserti128.VINSERTI128256ZERO{},
		&vinserti128.VINSERTI128256ONE{},
		// vextracti128
		&vextracti128.VEXTRACTI128256ZERO{},
		&vextracti128.VEXTRACTI128256ONE{},
		// vmovdqu
		vmovdqu.NewVMOVDQU128(),
		vmovdqu.NewVMOVDQU256(),
		vmovdqa.NewVMOVDQA128(),
		vmovdqa.NewVMOVDQA256(),
		// addps
		addps.NewADDPS128(),
		addps.NewVADDPS128(),
		addps.NewVADDPS256(),
		// mulps
		mulps.NewMULPS128(),
		mulps.NewVMULPS128(),
		mulps.NewVMULPS256(),
		// subps
		subps.NewSUBPS128(),
		subps.NewVSUBPS128(),
		subps.NewVSUBPS256(),
		// andps
		andps.NewANDPS128(),
		andps.NewVANDPS128(),
		andps.NewVANDPS256(),
		// orps
		orps.NewORPS128(),
		orps.NewVORPS128(),
		orps.NewVORPS256(),
		// xorps
		xorps.NewXORPS128(),
		xorps.NewVXORPS128(),
		xorps.NewVXORPS256(),
		// andnps
		andnps.NewANDNPS128(),
		andnps.NewVANDNPS128(),
		andnps.NewVANDNPS256(),
		// pmulhuw
		pmulhuw.NewPMULHUW128(),
		pmulhuw.NewVPMULHUW128(),
		pmulhuw.NewVPMULHUW256(),
		// addss
		addss.NewADDSS128(),
		addss.NewVADDSS128(),
		// divss
		divss.NewDIVSS128(),
		divss.NewVDIVSS128(),
		// divps
		divps.NewDIVPS128(),
		divps.NewVDIVPS128(),
		divps.NewVDIVPS256(),
		// subss
		subss.NewSUBSS128(),
		subss.NewVSUBSS128(),
		// cmpps (128-bit)
		cmpps.NewCMPPS128EQ(),
		cmpps.NewCMPPS128LT(),
		cmpps.NewCMPPS128LE(),
		cmpps.NewCMPPS128UNORD(),
		cmpps.NewCMPPS128NEQ(),
		cmpps.NewCMPPS128NLT(),
		cmpps.NewCMPPS128NLE(),
		cmpps.NewCMPPS128ORD(),
		// mulss
		mulss.NewMULSS128(),
		mulss.NewVMULSS128(),
		// movss
		movss.NewMOVSS128(),
		movss.NewVMOVSS128(),
		// cmpss
		cmpss.NewCMPSS128EQ(),
		cmpss.NewCMPSS128LT(),
		cmpss.NewCMPSS128LE(),
		cmpss.NewCMPSS128UNORD(),
		cmpss.NewCMPSS128NEQ(),
		cmpss.NewCMPSS128NLT(),
		cmpss.NewCMPSS128NLE(),
		cmpss.NewCMPSS128ORD(),
		cmpss.NewVCMPSS128EQ(),
		cmpss.NewVCMPSS128LT(),
		cmpss.NewVCMPSS128LE(),
		cmpss.NewVCMPSS128UNORD(),
		cmpss.NewVCMPSS128NEQ(),
		cmpss.NewVCMPSS128NLT(),
		cmpss.NewVCMPSS128NLE(),
		cmpss.NewVCMPSS128ORD(),
		// comiss
		comiss.NewCOMISS128(),
		comiss.NewVCOMISS128(),
		// ucomiss
		ucomiss.NewUCOMISS128(),
		ucomiss.NewVUCOMISS128(),
		// rcpps
		rcpps.NewRCPPS128(),
		rcpps.NewVRCPPS128(),
		rcpps.NewVRCPPS256(),
		// sqrtps
		sqrtps.NewSQRTPS128(),
		sqrtps.NewVSQRTPS128(),
		sqrtps.NewVSQRTPS256(),
		// rsqrtps
		rsqrtps.NewRSQRTPS128(),
		rsqrtps.NewVRSQRTPS128(),
		rsqrtps.NewVRSQRTPS256(),
		// rcpss
		rcpss.NewRCPSS128(),
		rcpss.NewVRCPSS128(),
		// rsqrtss
		rsqrtss.NewRSQRTSS128(),
		rsqrtss.NewVRSQRTSS128(),
		// sqrtss
		sqrtss.NewSQRTSS128(),
		sqrtss.NewVSQRTSS128(),
		// movaps
		movaps.NewMOVAPS128(),
		movaps.NewVMOVAPS128(),
		movaps.NewVMOVAPS256(),
		// movups
		movups.NewMOVUPS128(),
		movups.NewVMOVUPS128(),
		movups.NewVMOVUPS256(),
		// movlps
		movlps.NewMOVLPS64(),
		movlps.NewVMOVLPS64(),
		// movhpos
		movhps.NewMOVHPS64(),
		movhps.NewVMOVHPS64(),
		// unpckhps
		unpckhps.NewUNPCKHPS128(),
		unpckhps.NewVUNPCKHPS128(),
		unpckhps.NewVUNPCKHPS256(),
		// shufps
		shufps.NewSHUFPS128ZEROS(),
		shufps.NewSHUFPS128ONES(),
		shufps.NewSHUFPS128TWOS(),
		shufps.NewSHUFPS128THREES(),
		shufps.NewSHUFPS128MIXED(),
		shufps.NewSHUFPS128REVERSE(),
		shufps.NewVSHUFPS128ZEROS(),
		shufps.NewVSHUFPS128ONES(),
		shufps.NewVSHUFPS128TWOS(),
		shufps.NewVSHUFPS128THREES(),
		shufps.NewVSHUFPS128MIXED(),
		shufps.NewVSHUFPS128REVERSE(),
		shufps.NewVSHUFPS256ZEROS(),
		shufps.NewVSHUFPS256ONES(),
		shufps.NewVSHUFPS256TWOS(),
		shufps.NewVSHUFPS256THREES(),
		shufps.NewVSHUFPS256MIXED(),
		shufps.NewVSHUFPS256REVERSE(),
		// movlhps
		movlhps.NewMOVLHPS64(),
		movlhps.NewVMOVLHPS64(),
		// maxps
		maxps.NewMAXPS128(),
		maxps.NewVMAXPS128(),
		maxps.NewVMAXPS256(),
		// maxss
		maxss.NewMAXSS128(),
		maxss.NewVMAXSS128(),
		// minps
		minps.NewMINPS128(),
		minps.NewVMINPS128(),
		minps.NewVMINPS256(),
		// minss
		minss.NewMINSS128(),
		minss.NewVMINSS128(),
		// movmskps
		movmskps.NewMOVMSKPS128(),
		movmskps.NewVMOVMSKPS128(),
		movmskps.NewVMOVMSKPS256(),
		// movmskpd
		movmskpd.NewMOVMSKPD128(),
		movmskpd.NewVMOVMSKPD128(),
		movmskpd.NewVMOVMSKPD256(),
		// psadbw
		psadbw.NewPSADBW128(),
		psadbw.NewVPSADBW128(),
		psadbw.NewVPSADBW256(),
		// pavgb
		pavgb.NewPAVGB128(),
		pavgb.NewVPAVGB128(),
		pavgb.NewVPAVGB256(),
		// pavgw
		pavgw.NewPAVGW128(),
		pavgw.NewVPAVGW128(),
		pavgw.NewVPAVGW256(),
		// pmaxub
		pmaxub.NewPMAXUB128(),
		pmaxub.NewVPMAXUB128(),
		pmaxub.NewVPMAXUB256(),
		// pminub
		pminub.NewPMINUB128(),
		pminub.NewVPMINUB128(),
		pminub.NewVPMINUB256(),
		// pmaxsw
		pmaxsw.NewPMAXSW128(),
		pmaxsw.NewVPMAXSW128(),
		pmaxsw.NewVPMAXSW256(),
		// pminsw
		pminsw.NewPMINSW128(),
		pminsw.NewVPMINSW128(),
		pminsw.NewVPMINSW256(),
		// pextrw
		pextrw.NewPEXTRW128ZERO_IDX(),
		pextrw.NewPEXTRW128ONE_IDX(),
		pextrw.NewPEXTRW128TWO_IDX(),
		pextrw.NewPEXTRW128THREE_IDX(),
		pextrw.NewPEXTRW128FOUR_IDX(),
		pextrw.NewPEXTRW128FIVE_IDX(),
		pextrw.NewPEXTRW128SIX_IDX(),
		pextrw.NewPEXTRW128SEVEN_IDX(),
		pextrw.NewVPEXTRW128ZERO_IDX(),
		pextrw.NewVPEXTRW128ONE_IDX(),
		pextrw.NewVPEXTRW128TWO_IDX(),
		pextrw.NewVPEXTRW128THREE_IDX(),
		pextrw.NewVPEXTRW128FOUR_IDX(),
		pextrw.NewVPEXTRW128FIVE_IDX(),
		pextrw.NewVPEXTRW128SIX_IDX(),
		pextrw.NewVPEXTRW128SEVEN_IDX(),
		// pinsrw
		pinsrw.NewPINSRW128ZERO_IDX(),
		pinsrw.NewPINSRW128ONE_IDX(),
		pinsrw.NewPINSRW128TWO_IDX(),
		pinsrw.NewPINSRW128THREE_IDX(),
		pinsrw.NewPINSRW128FOUR_IDX(),
		pinsrw.NewPINSRW128FIVE_IDX(),
		pinsrw.NewPINSRW128SIX_IDX(),
		pinsrw.NewPINSRW128SEVEN_IDX(),
		pinsrw.NewVPINSRW128ZERO_IDX(),
		pinsrw.NewVPINSRW128ONE_IDX(),
		pinsrw.NewVPINSRW128TWO_IDX(),
		pinsrw.NewVPINSRW128THREE_IDX(),
		pinsrw.NewVPINSRW128FOUR_IDX(),
		pinsrw.NewVPINSRW128FIVE_IDX(),
		pinsrw.NewVPINSRW128SIX_IDX(),
		pinsrw.NewVPINSRW128SEVEN_IDX(),
		// pmovmskb
		pmovmskb.NewPMOVMSKB128(),
		pmovmskb.NewVPMOVMSKB128(),
		// vbroadcast
		vbroadcast.NewVBROADCASTSS128M32(),
		vbroadcast.NewVBROADCASTSS256M32(),
		vbroadcast.NewVBROADCASTSD256M64(),
		vbroadcast.NewVBROADCASTSF128256M128(),
		// vinsertf128
		vinsertf128.NewVINSERTF128256ZERO(),
		vinsertf128.NewVINSERTF128256ONE(),
		// vextractf128
		vextractf128.NewVEXTRACTF128256ZERO(),
		vextractf128.NewVEXTRACTF128256ONE(),
		// vmaskmov
		vmaskmov.NewVMASKMOVPS128LOAD(),
		vmaskmov.NewVMASKMOVPD128LOAD(),
		vmaskmov.NewVMASKMOVPD256LOAD(),
		vmaskmov.NewVMASKMOVPS128STORE(),
		vmaskmov.NewVMASKMOVPD128STORE(),
		vmaskmov.NewVMASKMOVPD256STORE(),
		// vpermilps
		vpermilps.NewVPERMILPS128(),
		vpermilps.NewVPERMILPS128IDENTITY(),
		vpermilps.NewVPERMILPS128ALL_ZERO(),
		vpermilps.NewVPERMILPS128ALL_ONE(),
		vpermilps.NewVPERMILPS128ALL_TWO(),
		vpermilps.NewVPERMILPS128ALL_THREE(),
		vpermilps.NewVPERMILPS128REVERSE(),
		vpermilps.NewVPERMILPS256(),
		vpermilps.NewVPERMILPS256IDENTITY(),
		vpermilps.NewVPERMILPS256ALL_ZERO(),
		vpermilps.NewVPERMILPS256ALL_ONE(),
		vpermilps.NewVPERMILPS256ALL_TWO(),
		vpermilps.NewVPERMILPS256ALL_THREE(),
		vpermilps.NewVPERMILPS256REVERSE(),
		// vpermilpd
		vpermilpd.NewVPERMILPD128(),
		vpermilpd.NewVPERMILPD128IDENTITY(),
		vpermilpd.NewVPERMILPD128ALL_ZERO(),
		vpermilpd.NewVPERMILPD128ALL_ONE(),
		vpermilpd.NewVPERMILPD128REVERSE(),
		vpermilpd.NewVPERMILPD256(),
		vpermilpd.NewVPERMILPD256IDENTITY(),
		vpermilpd.NewVPERMILPD256ALL_ZERO(),
		vpermilpd.NewVPERMILPD256ALL_ONE(),
		vpermilpd.NewVPERMILPD256REVERSE(),
		// vperm2f128
		vperm2f128.NewVPERM2F128256LOWA_HIGHA(),
		vperm2f128.NewVPERM2F128256LOWB_HIGHB(),
		vperm2f128.NewVPERM2F128256LOWA_HIGHB(),
		vperm2f128.NewVPERM2F128256LOWB_HIGHA(),
		vperm2f128.NewVPERM2F128256LOWA_LOWA(),
		vperm2f128.NewVPERM2F128256HIGHA_HIGHA(),
		vperm2f128.NewVPERM2F128256LOWB_LOWB(),
		vperm2f128.NewVPERM2F128256HIGHB_HIGHB(),
		vperm2f128.NewVPERM2F128256ZEROED_HIGHB(),
		vperm2f128.NewVPERM2F128256LOWA_ZEROED(),
		vperm2f128.NewVPERM2F128256ZEROED_ZEROED(),
		// vpblendd (128-bit)
		vpblendd.NewVPBLENDD128NONE(),
		vpblendd.NewVPBLENDD128ALL(),
		vpblendd.NewVPBLENDD128LOW_HALF(),
		vpblendd.NewVPBLENDD128HIGH_HALF(),
		vpblendd.NewVPBLENDD128EVEN(),
		vpblendd.NewVPBLENDD128ODD(),
		// vperm2i128
		vperm2i128.NewVPERM2I128256LOWA_HIGHA(),
		vperm2i128.NewVPERM2I128256LOWB_HIGHB(),
		vperm2i128.NewVPERM2I128256LOWA_HIGHB(),
		vperm2i128.NewVPERM2I128256LOWB_HIGHA(),
		vperm2i128.NewVPERM2I128256LOWA_LOWA(),
		vperm2i128.NewVPERM2I128256HIGHA_HIGHA(),
		vperm2i128.NewVPERM2I128256LOWB_LOWB(),
		vperm2i128.NewVPERM2I128256HIGHB_HIGHB(),
		vperm2i128.NewVPERM2I128256ZEROED_HIGHB(),
		vperm2i128.NewVPERM2I128256LOWA_ZEROED(),
		vperm2i128.NewVPERM2I128256ZEROED_ZEROED(),
		// vpermps
		vpermps.NewVPERMPS256(),
		// vpermd
		vpermd.NewVPERMD256(),
		// vpermq
		vpermq.NewVPERMQ256IDENTITY(),
		vpermq.NewVPERMQ256REVERSE(),
		vpermq.NewVPERMQ256ALL_ZEROS(),
		vpermq.NewVPERMQ256ALL_ONES(),
		vpermq.NewVPERMQ256ALL_TWOS(),
		vpermq.NewVPERMQ256ALL_THREES(),
		// vtestps
		vtestps.NewVTESTPS128(),
		vtestps.NewVTESTPS256(),
		// vtestpd
		vtestpd.NewVTESTPD128(),
		vtestpd.NewVTESTPD256(),
		// vgatherdpd
		vgatherdpd.NewVGATHERDPD128(),
		vgatherdpd.NewVGATHERDPD256(),
		// vgatherqpd
		vgatherqdp.NewVGATHERQPD128(),
		vgatherqdp.NewVGATHERQPD256(),
		// vpgatherdd
		vpgatherdd.NewVPGATHERDD128(),
		vpgatherdd.NewVPGATHERDD256(),
		// vpgatherqd
		vpgatherqd.NewVPGATHERQD128(),
		vpgatherqd.NewVPGATHERQD256(),
		// vpgatherdq
		vpgatherdq.NewVPGATHERDQ128(),
		vpgatherdq.NewVPGATHERDQ256(),
		// vpgatherqq
		vpgatherqq.NewVPGATHERQQ128(),
		vpgatherqq.NewVPGATHERQQ256(),
		// vpmaskmovd/q
		vpmaskmov.NewVPMASKMOVD128(),
		vpmaskmov.NewVPMASKMOVD256(),
		vpmaskmov.NewVPMASKMOVQ128(),
		vpmaskmov.NewVPMASKMOVQ256(),
		// vpsllvd
		vpsllvd.NewVPSLLVD128(),
		vpsllvd.NewVPSLLVD256(),
		// vpsllvq
		vpsllvq.NewVPSLLVQ128(),
		vpsllvq.NewVPSLLVQ256(),
		// vpsrlvd
		vpsrlvd.NewVPSRLVD128(),
		vpsrlvd.NewVPSRLVD256(),
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
