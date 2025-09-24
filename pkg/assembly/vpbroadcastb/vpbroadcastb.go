package vpbroadcastb

type VPBROADCASTB struct {
}

func (v *VPBROADCASTB) InputSize() int {
	return 8
}

func (v *VPBROADCASTB) OutputSize() int {
	return 256
}

func (v *VPBROADCASTB) Name() string {
	return "VPBROADCASTB"
}

func (v *VPBROADCASTB) Description() string {
	return "TODO"
}

func (v *VPBROADCASTB) Stub() string {
	// Find a way to automatically include the generated assembly here
	return "TODO"
}

func (v *VPBROADCASTB) Assembly() string {
	// Find a way to automatically include the generated assembly here
	return "TODO"
}

func (v *VPBROADCASTB) Run(input []byte) (output []byte) {
	ret := [32]byte{}
	vpbroadcastb(input[0], &ret)
	return ret[:]
}
