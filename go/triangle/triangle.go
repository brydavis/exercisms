package triangle

import "sort"

type Kind struct{}

var Iso, Equ, Sca, NaT Kind

func KindFromSides(a, b, c float64) Kind {
	triangle := []float64{a, b, c}
	sort.Sort(sort.Reverse(sort.Float64Slice(triangle)))

	max := triangle[0]
	min1 := triangle[1]
	min2 := triangle[2]

	if (min1 + min2) > max {
		switch {
		case max == min1 && max != min2:
			return Iso
		case max == min1 && max == min2:
			return Equ
		case max != min1 && max != min2 && min1 != min2:
			return Sca
		}
	}

	return NaT
}
