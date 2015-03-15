package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	product := 1
	d := len(digits)
	if d < span {
		return 0, errors.New("Span is larger than number of digits")
	} else if d > 0 {
		for i := 0; i < d; i++ {
			if i+span <= d {
				p := 1
				for _, v := range digits[i : i+span] {
					vv, _ := strconv.Atoi(string(v))
					p *= vv
				}
				if p > product {
					product = p
				}
			}
		}
	}
	return product, nil
}
