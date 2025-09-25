package vpbroadcastb

type Vpbroadcastb512 struct {
}

func (v *Vpbroadcastb512) InputSize() int {
	return 8
}

func (v *Vpbroadcastb512) OutputSize() int {
	return 512
}

func (v *Vpbroadcastb512) Name() string {
	return "VPBROADCASTB ZMM (512 bit)"
}

func (v *Vpbroadcastb512) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb512) Stub() string {
	// Find a way to automatically include the generated assembly here
	return "TODO"
}

func (v *Vpbroadcastb512) Assembly() string {
	// Find a way to automatically include the generated assembly here
	return "TODO"
}

func (v *Vpbroadcastb512) Run(input []byte) (output []byte) {
	ret := [64]byte{}
	vpbroadcastb512(input[0], &ret)
	return ret[:]
}
