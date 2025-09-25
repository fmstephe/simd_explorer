package all

import (
	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/assembly/vpbroadcastb"
)

func Instructions() []assembly.Instruction {
	return []assembly.Instruction{
		&vpbroadcastb.Vpbroadcastb128{},
		&vpbroadcastb.Vpbroadcastb256{},
		&vpbroadcastb.Vpbroadcastb512{},
	}
}
