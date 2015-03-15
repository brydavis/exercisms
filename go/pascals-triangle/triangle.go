package pascal

func Triangle(rows int) (store [][]int) {
	for n := 0; n <= rows-1; n++ {
		var r []int
		for i := 0; i <= n; i++ {
			r = append(r, Combination(n, i))
		}
		store = append(store, r)
	}
	return store
}

func Factorial(n int) int {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}

func Combination(n, r int) int {
	return Factorial(n) / (Factorial(r) * Factorial(n-r))
}
