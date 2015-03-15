package pythagorean

type Triplet [3]int

// func Range(min, max int) (result []Triplet) {
// 	for a := min; a <= max; a++ {
// 		for b := a; b <= max; b++ {
// 			aSq, bSq := a*a, b*b
// 			c := math.Sqrt(float64(aSq) + float64(bSq))
// 			if math.Mod(c, 1) == 0 && int(c) <= max {
// 				result = append(result, Triplet{a, b, int(c)})
// 			}
// 		}
// 	}
// 	return
// }

// func Range(min, max int) []Triplet {
// 	res := []Triplet{}
// 	for a := min; a <= max; a++ {
// 		for b := a; b <= max; b++ {
// 			for c := b; c <= max; c++ {
// 				if a*a+b*b == c*c {
// 					res = append(res, Triplet{a, b, c})
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// func Range(min, max int) (result []Triplet) {
// 	for a := min; a <= max; a++ {
// 		for b := a; b <= max; b++ {
// 			for c := b; c <= max; c++ {
// 				if a*a+b*b == c*c {
// 					result = append(result, Triplet{a, b, c})
// 				}
// 			}
// 		}
// 	}
// 	return
// }

func Sum(p int) (result []Triplet) {
	for a := 1; a <= p; a++ {
		for b := a; b <= p; b++ {
			for c := b; c <= p; c++ {
				if a*a+b*b == c*c && (a+b+c) == p {
					result = append(result, Triplet{a, b, int(c)})
				}
			}
		}
	}
	return
}

func Range(min, max int) (result []Triplet) {
	for a := min; a <= max; a++ {
		go func() {
			for b := a; b <= max; b++ {
				for c := b; c <= max; c++ {
					if a*a+b*b == c*c {
						result = append(result, Triplet{a, b, c})
					}
				}
			}
		}()
	}
	return
}

// func Brute(min, max, deep int) {
// 	if deep > 0 {
// 		for a := min; a <= max; a++ {
// 			Brute(a, max, deep-1)
// 		}
// 	} else if a*a+b*b == c*c {
// 		result = append(result, Triplet{a, b, c})
// 	}
// }
