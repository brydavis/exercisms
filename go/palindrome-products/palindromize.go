package palindrome // Code Passes Test, Yet Not Optimized

import (
	"errors"
	"math"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Reverse(text string) (result string) {
	for _, v := range text {
		result = string(v) + result
	}
	return
}

func Products(fmin, fmax int) (Product, Product, error) {
	pmin := Product{}
	pmax := Product{}

	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax...")
	}

	for x := fmin; x <= fmax; x++ {
		for y := x; y <= fmax; y++ {
			z := strconv.Itoa(x * y)

			var half int
			var left string
			// var middle string
			var right string

			if math.Mod(float64(len(z)), 2.0) == 0.0 {
				half = len(z) / 2
				left = z[:half]
				right = z[half:]
			} else {
				half = (len(z) - 1) / 2
				left = z[:half]
				// middle = z[half : half+1]
				right = z[half+1:]
			}

			if left == Reverse(right) {
				if pmin.Product == 0 {
					pmin.Product = x * y
					pmin.Factorizations = append(pmin.Factorizations, [2]int{x, y})
				}

				if pmax.Product <= x*y {
					pmaxTemp := Product{Product: x * y}
					pmax.Product = x * y
					for _, v := range pmax.Factorizations {
						if (v[0] * v[1]) == pmax.Product {
							pmaxTemp.Factorizations = append(pmaxTemp.Factorizations, v)
						}
					}
					pmax = pmaxTemp
					pmax.Factorizations = append(pmax.Factorizations, [2]int{x, y})
				}

			}
			// fmt.Println(x, y, z, l, "left("+left+")", "middle("+middle+")", "right("+right+")", left == Reverse(right)) // "half("+half+")",
		}
	}
	if pmin.Product == pmax.Product && pmin.Product > 1 {
		pmin = Product{}
	}
	// fmt.Println(pmin, pmax)
	if pmax.Product == 0 && len(pmax.Factorizations) == 0 {
		return pmin, pmax, errors.New("No palindromes...")
	}

	return pmin, pmax, nil
}
