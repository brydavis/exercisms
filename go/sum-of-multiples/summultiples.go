package summultiples

func MultipleSummer(divs ...int) func(limit int) int {
	return func(limit int) int {
		total := 0
		for i := 1; i < limit; i++ {
			for _, v := range divs {
				if i%v == 0 {
					total += i
					break
				}
			}
		}
		return total
	}
}
