package grains

func Square(input int) (val uint64, err bool) {
	if input < 1 || input > 64 {
		return 0, true
	}

	val = 1
	for i := 1; i < input; i++ {
		val += val
	}
	return val, err
}

func Total() (val uint64) {
	for i := 1; i < 65; i++ {
		v, _ := Square(i)
		val += v
	}
	return val
}
