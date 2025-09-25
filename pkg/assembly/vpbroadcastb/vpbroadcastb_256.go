package vpbroadcastb

type Vpbroadcastb256 struct {
}

func (v *Vpbroadcastb256) InputSize() int {
	return 8
}

func (v *Vpbroadcastb256) OutputSize() int {
	return 256
}

func (v *Vpbroadcastb256) Name() string {
	return "VPBROADCASTB YMM (256 bit)"
}

func (v *Vpbroadcastb256) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb256) Stub() string {
	// Find a way to automatically include the generated assembly here
	return "TODO"
}

func (v *Vpbroadcastb256) Assembly() string {
	// Find a way to automatically include the generated assembly here
	return "TODO"
}

func (v *Vpbroadcastb256) Run(input []byte) (output []byte) {
	ret := [32]byte{}
	vpbroadcastb256(input[0], &ret)
	return ret[:]
}
