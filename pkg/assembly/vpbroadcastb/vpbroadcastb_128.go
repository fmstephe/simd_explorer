package vpbroadcastb

type Vpbroadcastb128 struct {
}

func (v *Vpbroadcastb128) InputSize() int {
	return 8
}

func (v *Vpbroadcastb128) OutputSize() int {
	return 128
}

func (v *Vpbroadcastb128) Name() string {
	return "VPBROADCASTB XMM (128 bit)"
}

func (v *Vpbroadcastb128) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb128) Stub() string {
	// Find a way to automatically include the generated assembly here
	return "TODO"
}

func (v *Vpbroadcastb128) Assembly() string {
	// Find a way to automatically include the generated assembly here
	return "TODO"
}

func (v *Vpbroadcastb128) Run(input []byte) (output []byte) {
	ret := [16]byte{}
	vpbroadcastb128(input[0], &ret)
	return ret[:]
}
