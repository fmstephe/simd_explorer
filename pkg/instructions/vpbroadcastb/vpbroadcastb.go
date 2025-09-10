package vpbroadcastb

func Run(b byte) []byte {
	ret := [8]byte{}
	vpbroadcastb(b, &ret)
	return ret[:]
}
