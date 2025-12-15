package all

import (
	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/assembly/addps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/addss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/andnps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/andps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/blendpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/blendps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/blendvpb"
	"github.com/fmstephe/simd_explorer/pkg/assembly/blendvpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/blendvps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/blendvpw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cmpps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cmpss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/comiss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cvtdq2pd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cvtdq2ps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cvtpd2dq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cvtpd2ps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cvtps2dq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/cvtps2pd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/divps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/divss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/dppd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/dpps"
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
	"github.com/fmstephe/simd_explorer/pkg/assembly/packss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/packus"
	"github.com/fmstephe/simd_explorer/pkg/assembly/padd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/palignr"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pand"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pandn"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pavgb"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pavgw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pcmpeq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pcmpgt"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pextrw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pinsrw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmaddhw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmaddubsw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmaxs"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmaxu"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmins"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pminu"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmovmskb"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmulhuw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmulhw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmull"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pmuludq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/por"
	"github.com/fmstephe/simd_explorer/pkg/assembly/psadbw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pshufb"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pshufd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pshufhw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pshuflw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/psll"
	"github.com/fmstephe/simd_explorer/pkg/assembly/psra"
	"github.com/fmstephe/simd_explorer/pkg/assembly/psrl"
	"github.com/fmstephe/simd_explorer/pkg/assembly/psub"
	"github.com/fmstephe/simd_explorer/pkg/assembly/punpckh"
	"github.com/fmstephe/simd_explorer/pkg/assembly/punpckl"
	"github.com/fmstephe/simd_explorer/pkg/assembly/pxor"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rcpps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rcpss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/roundp"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rounds"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rsqrtps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/rsqrtss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/shufps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/sqrtps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/sqrtss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/subps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/subss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/ucomiss"
	"github.com/fmstephe/simd_explorer/pkg/assembly/unpckhps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vaddsubpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vaddsubps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vbroadcast"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vbroadcasti128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vextractf128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vextracti128"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vgatherdpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vgatherqdp"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vhaddpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vhaddps"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vhsubpd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vhsubps"
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
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpmaddwd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpmaskmov"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpmovsx"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpmovzx"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpmulhrsw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpsllvd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpsllvq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpsllvw"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpsravd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpsrlvd"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpsrlvq"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vptest"
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
		// dppd (dot product packed doubles, imm8 variants)
		dppd.NewVDPPD128NONE(),
		dppd.NewVDPPD128LOW_ONLY(),
		dppd.NewVDPPD128HIGH_ONLY(),
		dppd.NewVDPPD128ALL(),
		// dpps (dot product packed singles, imm8 variants)
		dpps.NewVDPPS128NONE(),
		dpps.NewVDPPS128LOW_ONLY(),
		dpps.NewVDPPS128HIGH_ONLY(),
		dpps.NewVDPPS128ALL(),
		// roundp (round packed singles/doubles, imm8 variants)
		roundp.NewVROUNDPS128ZERO(),
		roundp.NewVROUNDPS128ONE(),
		roundp.NewVROUNDPS128TWO(),
		roundp.NewVROUNDPS128THREE(),
		roundp.NewVROUNDPS256ZERO(),
		roundp.NewVROUNDPS256ONE(),
		roundp.NewVROUNDPS256TWO(),
		roundp.NewVROUNDPS256THREE(),
		roundp.NewVROUNDPD128ZERO(),
		roundp.NewVROUNDPD128ONE(),
		roundp.NewVROUNDPD128TWO(),
		roundp.NewVROUNDPD128THREE(),
		roundp.NewVROUNDPD256ZERO(),
		roundp.NewVROUNDPD256ONE(),
		roundp.NewVROUNDPD256TWO(),
		roundp.NewVROUNDPD256THREE(),
		// rounds (scalar round ss/sd, with base + vals)
		rounds.NewVROUNDSS128ZERO(),
		rounds.NewVROUNDSS128ONE(),
		rounds.NewVROUNDSS128TWO(),
		rounds.NewVROUNDSS128THREE(),
		rounds.NewVROUNDSD128ZERO(),
		rounds.NewVROUNDSD128ONE(),
		rounds.NewVROUNDSD128TWO(),
		rounds.NewVROUNDSD128THREE(),
		// converts (packed)
		cvtps2pd.NewVCVTPS2PD128(),
		cvtps2pd.NewVCVTPS2PD256(),
		cvtpd2ps.NewVCVTPD2PS128(),
		cvtpd2ps.NewVCVTPD2PS256(),
		cvtdq2ps.NewVCVTDQ2PS128(),
		cvtdq2ps.NewVCVTDQ2PS256(),
		cvtps2dq.NewVCVTPS2DQ128(),
		cvtps2dq.NewVCVTPS2DQ256(),
		cvtpd2dq.NewVCVTPD2DQ128(),
		cvtpd2dq.NewVCVTPD2DQ256(),
		cvtdq2pd.NewVCVTDQ2PD128(),
		cvtdq2pd.NewVCVTDQ2PD256(),
		// pshufb
		pshufb.NewVPSHUFB128(),
		pshufb.NewVPSHUFB256(),
		// pshufd (imm8 variants)
		pshufd.NewVPSHUFD128IDENTITY(),
		pshufd.NewVPSHUFD128REVERSE(),
		pshufd.NewVPSHUFD128EVENS(),
		pshufd.NewVPSHUFD128ODDS(),
		pshufd.NewVPSHUFD256IDENTITY(),
		pshufd.NewVPSHUFD256REVERSE(),
		pshufd.NewVPSHUFD256EVENS(),
		pshufd.NewVPSHUFD256ODDS(),
		// pshufhw (imm8 variants on high words)
		pshufhw.NewVPSHUFHW128IDENTITY(),
		pshufhw.NewVPSHUFHW128REVERSE(),
		pshufhw.NewVPSHUFHW128EVENS(),
		pshufhw.NewVPSHUFHW128ODDS(),
		pshufhw.NewVPSHUFHW256IDENTITY(),
		pshufhw.NewVPSHUFHW256REVERSE(),
		pshufhw.NewVPSHUFHW256EVENS(),
		pshufhw.NewVPSHUFHW256ODDS(),
		// pshuflw (imm8 variants on low words)
		pshuflw.NewVPSHUFLW128IDENTITY(),
		pshuflw.NewVPSHUFLW128REVERSE(),
		pshuflw.NewVPSHUFLW128EVENS(),
		pshuflw.NewVPSHUFLW128ODDS(),
		pshuflw.NewVPSHUFLW256IDENTITY(),
		pshuflw.NewVPSHUFLW256REVERSE(),
		pshuflw.NewVPSHUFLW256EVENS(),
		pshuflw.NewVPSHUFLW256ODDS(),
		// palignr imm8=17 examples
		palignr.NewVPALIGNR128SEVENTEEN(),
		palignr.NewVPALIGNR256SEVENTEEN(),
		// palignr
		palignr.NewVPALIGNR128ZERO(),
		palignr.NewVPALIGNR128ONE(),
		palignr.NewVPALIGNR128TWO(),
		palignr.NewVPALIGNR128THREE(),
		palignr.NewVPALIGNR128FOUR(),
		palignr.NewVPALIGNR128SIXTEEN(),
		palignr.NewVPALIGNR128THIRTYTWO(),
		palignr.NewVPALIGNR128THIRTYTHREE(),
		palignr.NewVPALIGNR256ZERO(),
		palignr.NewVPALIGNR256ONE(),
		palignr.NewVPALIGNR256TWO(),
		palignr.NewVPALIGNR256THREE(),
		palignr.NewVPALIGNR256EIGHT(),
		palignr.NewVPALIGNR256FOUR(),
		palignr.NewVPALIGNR256SIXTEEN(),
		palignr.NewVPALIGNR256THIRTYTWO(),
		palignr.NewVPALIGNR256THIRTYTHREE(),
		// vpor
		por.NewVPOR128(),
		por.NewVPOR256(),
		// vpandn
		pandn.NewVPANDN128(),
		pandn.NewVPANDN256(),
		// vpxor
		pxor.NewVPXOR128(),
		pxor.NewVPXOR256(),
		// pcmpeq
		pcmpeq.NewVPCMPEQB128(),
		pcmpeq.NewVPCMPEQB256(),
		pcmpeq.NewVPCMPEQW128(),
		pcmpeq.NewVPCMPEQW256(),
		pcmpeq.NewVPCMPEQD128(),
		pcmpeq.NewVPCMPEQD256(),
		pcmpeq.NewVPCMPEQQ128(),
		pcmpeq.NewVPCMPEQQ256(),
		// pcmpgt
		pcmpgt.NewVPCMPGTB128(),
		pcmpgt.NewVPCMPGTB256(),
		pcmpgt.NewVPCMPGTW128(),
		pcmpgt.NewVPCMPGTW256(),
		pcmpgt.NewVPCMPGTD128(),
		pcmpgt.NewVPCMPGTD256(),
		pcmpgt.NewVPCMPGTQ128(),
		pcmpgt.NewVPCMPGTQ256(),
		// pand (bitwise AND on integers)
		pand.NewVPAND128(),
		pand.NewVPAND256(),
		// pmaddubsw
		pmaddubsw.NewVPMADDUBSW128(),
		pmaddubsw.NewVPMADDUBSW256(),
		// pmuludq (unsigned dword multiply to quadword)
		pmuludq.NewVPMULUDQ128(),
		pmuludq.NewVPMULUDQ256(),
		// padd (AVX/AVX2 integer adds)
		padd.NewVPADDB128(),
		padd.NewVPADDB256(),
		padd.NewVPADDW128(),
		padd.NewVPADDW256(),
		padd.NewVPADDD128(),
		padd.NewVPADDD256(),
		padd.NewVPADDQ128(),
		padd.NewVPADDQ256(),
		// pmull (AVX/AVX2/AVX-512 integer low multiplies)
		pmull.NewVPMULLW128(),
		pmull.NewVPMULLW256(),
		pmull.NewVPMULLD128(),
		pmull.NewVPMULLD256(),
		pmull.NewVPMULLQ128(),
		pmull.NewVPMULLQ256(),
		// pmaddhw (AVX/AVX2: pairwise multiply words and add to 32-bit)
		pmaddhw.NewVPMADDHW128(),
		pmaddhw.NewVPMADDHW256(),
		// vpmaddwd (AVX/AVX2: multiply adjacent signed words and add into dwords)
		vpmaddwd.NewVPMADDWD128(),
		vpmaddwd.NewVPMADDWD256(),
		// pmulhw (AVX/AVX2 signed high multiply words)
		pmulhw.NewVPMULHW128(),
		pmulhw.NewVPMULHW256(),
		// vpmulhrsw (AVX/AVX2: multiply high with round and scale, signed 16→16)
		vpmulhrsw.NewVPMULHRSW128(),
		vpmulhrsw.NewVPMULHRSW256(),
		// psub (AVX/AVX2 integer subs)
		psub.NewVPSUBB128(),
		psub.NewVPSUBB256(),
		psub.NewVPSUBW128(),
		psub.NewVPSUBW256(),
		psub.NewVPSUBD128(),
		psub.NewVPSUBD256(),
		psub.NewVPSUBQ128(),
		psub.NewVPSUBQ256(),
		// pminu (unsigned mins)
		pminu.NewVPMINUB128(),
		pminu.NewVPMINUB256(),
		pminu.NewVPMINUW128(),
		pminu.NewVPMINUW256(),
		pminu.NewVPMINUD128(),
		pminu.NewVPMINUD256(),
		// psll (variable, register-count)
		psll.NewVPSLLW128(),
		psll.NewVPSLLW256(),
		psll.NewVPSLLD128(),
		psll.NewVPSLLD256(),
		psll.NewVPSLLQ128(),
		psll.NewVPSLLQ256(),
		// packss
		packss.NewVPACKSSWB128(),
		packss.NewVPACKSSWB256(),
		packss.NewVPACKSSDW128(),
		packss.NewVPACKSSDW256(),
		// packus
		packus.NewVPACKUSWB128(),
		packus.NewVPACKUSWB256(),
		packus.NewVPACKUSDW128(),
		packus.NewVPACKUSDW256(),
		// psrl (variable, register-count)
		psrl.NewVPSRLW128(),
		psrl.NewVPSRLW256(),
		psrl.NewVPSRLD128(),
		psrl.NewVPSRLD256(),
		psrl.NewVPSRLQ128(),
		psrl.NewVPSRLQ256(),
		// psra (variable, register-count)
		psra.NewVPSRAW128(),
		psra.NewVPSRAW256(),
		psra.NewVPSRAD128(),
		psra.NewVPSRAD256(),
		psra.NewVPSRAQ128(),
		psra.NewVPSRAQ256(),
		// punpckl (unpack low interleaves)
		punpckl.NewVPUNPCKLBW128(),
		punpckl.NewVPUNPCKLBW256(),
		punpckl.NewVPUNPCKLWD128(),
		punpckl.NewVPUNPCKLWD256(),
		punpckl.NewVPUNPCKLDQ128(),
		punpckl.NewVPUNPCKLDQ256(),
		punpckl.NewVPUNPCKLQDQ128(),
		punpckl.NewVPUNPCKLQDQ256(),
		// punpckh (unpack high interleaves)
		punpckh.NewVPUNPCKHBW128(),
		punpckh.NewVPUNPCKHBW256(),
		punpckh.NewVPUNPCKHWD128(),
		punpckh.NewVPUNPCKHWD256(),
		punpckh.NewVPUNPCKHDQ128(),
		punpckh.NewVPUNPCKHDQ256(),
		punpckh.NewVPUNPCKHQDQ128(),
		punpckh.NewVPUNPCKHQDQ256(),
		// pmins (signed mins)
		pmins.NewVPMINSB128(),
		pmins.NewVPMINSB256(),
		pmins.NewVPMINSW128(),
		pmins.NewVPMINSW256(),
		pmins.NewVPMINSD128(),
		pmins.NewVPMINSD256(),
		// pmaxs (signed maxs)
		pmaxs.NewVPMAXSB128(),
		pmaxs.NewVPMAXSB256(),
		pmaxs.NewVPMAXSW128(),
		pmaxs.NewVPMAXSW256(),
		pmaxs.NewVPMAXSD128(),
		pmaxs.NewVPMAXSD256(),
		// pmaxu (unsigned maxs)
		pmaxu.NewVPMAXUB128(),
		pmaxu.NewVPMAXUB256(),
		pmaxu.NewVPMAXUW128(),
		pmaxu.NewVPMAXUW256(),
		pmaxu.NewVPMAXUD128(),
		pmaxu.NewVPMAXUD256(),
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
		// vblendps (128-bit)
		blendps.NewVBLENDPS128NONE(),
		blendps.NewVBLENDPS128ALL(),
		blendps.NewVBLENDPS128LOW_HALF(),
		blendps.NewVBLENDPS128HIGH_HALF(),
		blendps.NewVBLENDPS128EVEN(),
		blendps.NewVBLENDPS128ODD(),
		// vblendps (256-bit)
		blendps.NewVBLENDPS256NONE(),
		blendps.NewVBLENDPS256ALL(),
		blendps.NewVBLENDPS256LOW_HALF(),
		blendps.NewVBLENDPS256HIGH_HALF(),
		blendps.NewVBLENDPS256EVEN(),
		blendps.NewVBLENDPS256ODD(),
		// vblendpd (128-bit)
		blendpd.NewVBLENDPD128NONE(),
		blendpd.NewVBLENDPD128ALL(),
		blendpd.NewVBLENDPD128LOW_HALF(),
		blendpd.NewVBLENDPD128HIGH_HALF(),
		blendpd.NewVBLENDPD128EVEN(),
		blendpd.NewVBLENDPD128ODD(),
		// vblendpd (256-bit)
		blendpd.NewVBLENDPD256NONE(),
		blendpd.NewVBLENDPD256ALL(),
		blendpd.NewVBLENDPD256LOW_HALF(),
		blendpd.NewVBLENDPD256HIGH_HALF(),
		blendpd.NewVBLENDPD256EVEN(),
		blendpd.NewVBLENDPD256ODD(),
		// vblendvps
		blendvps.NewVBLENDVPS128(),
		blendvps.NewVBLENDVPS256(),
		// vblendvpd
		blendvpd.NewVBLENDVPD128(),
		blendvpd.NewVBLENDVPD256(),
		// vblendvpb
		blendvpb.NewVBLENDVPB128(),
		blendvpb.NewVBLENDVPB256(),
		// vblendvpw
		blendvpw.NewVBLENDVPW128(),
		blendvpw.NewVBLENDVPW256(),
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
		// vptest (integer)
		vptest.NewVPTEST128(),
		vptest.NewVPTEST256(),
		// vhaddps/pd
		vhaddps.NewVHADDPS128(),
		vhaddps.NewVHADDPS256(),
		vhaddpd.NewVHADDPD128(),
		vhaddpd.NewVHADDPD256(),
		// vhsubps/pd
		vhsubps.NewVHSUBPS128(),
		vhsubps.NewVHSUBPS256(),
		vhsubpd.NewVHSUBPD128(),
		vhsubpd.NewVHSUBPD256(),
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
		// vpsllvw
		vpsllvw.NewVPSLLVW128(),
		vpsllvw.NewVPSLLVW256(),
		// vpsllvw
		vpsllvd.NewVPSLLVD128(),
		vpsllvd.NewVPSLLVD256(),
		// vpsllvq
		vpsllvq.NewVPSLLVQ128(),
		vpsllvq.NewVPSLLVQ256(),
		// vpsrlvd
		vpsrlvd.NewVPSRLVD128(),
		vpsrlvd.NewVPSRLVD256(),
		// vpsrlvq
		vpsrlvq.NewVPSRLVQ128(),
		vpsrlvq.NewVPSRLVQ256(),
		// vpsravd
		vpsravd.NewVPSRAVD128(),
		vpsravd.NewVPSRAVD256(),
		// vpmovsx* (sign-extend)
		vpmovsx.NewVPMOVSXBW128(),
		vpmovsx.NewVPMOVSXBW256(),
		vpmovsx.NewVPMOVSXBD128(),
		vpmovsx.NewVPMOVSXBD256(),
		vpmovsx.NewVPMOVSXBQ128(),
		vpmovsx.NewVPMOVSXBQ256(),
		vpmovsx.NewVPMOVSXWD128(),
		vpmovsx.NewVPMOVSXWD256(),
		vpmovsx.NewVPMOVSXWQ128(),
		vpmovsx.NewVPMOVSXWQ256(),
		vpmovsx.NewVPMOVSXDQ128(),
		vpmovsx.NewVPMOVSXDQ256(),
		// vpmovzx* (zero-extend)
		vpmovzx.NewVPMOVZXBW128(),
		vpmovzx.NewVPMOVZXBW256(),
		vpmovzx.NewVPMOVZXBD128(),
		vpmovzx.NewVPMOVZXBD256(),
		vpmovzx.NewVPMOVZXBQ128(),
		vpmovzx.NewVPMOVZXBQ256(),
		vpmovzx.NewVPMOVZXWD128(),
		vpmovzx.NewVPMOVZXWD256(),
		vpmovzx.NewVPMOVZXWQ128(),
		vpmovzx.NewVPMOVZXWQ256(),
		vpmovzx.NewVPMOVZXDQ128(),
		vpmovzx.NewVPMOVZXDQ256(),
		// vaddsub*
		vaddsubps.NewVADDSUBPS128(),
		vaddsubps.NewVADDSUBPS256(),
		vaddsubpd.NewVADDSUBPD128(),
		vaddsubpd.NewVADDSUBPD256(),
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
