package luhn

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Valid(n string) (ok bool) {
	n = strings.Replace(n, " ", "", -1)
	if len(n) != 0 {
		n = Reverse(n)
		var (
			result string
			total  int
		)
		for k, v := range n {
			if math.Mod(float64(k), 2) != 0 {
				vv, err := strconv.Atoi(string(v))
				check(err)
				vvv := vv * 2
				if vvv > 9 {
					vvv = vvv - 9
				}
				result += strconv.Itoa(vvv)
			} else {
				result += string(v)
			}
		}
		for _, v := range result {
			vv, err := strconv.Atoi(string(v))
			check(err)
			total += vv
		}
		if rem := math.Mod(float64(total), 10); rem == 0 {
			ok = true
		}
	}
	return ok
}

func AddCheck(raw string) (luhn string) {
	if ok := Valid(raw); !ok {
		for i := 0; i <= 9; i++ {
			n := strconv.Itoa(i)
			rawn := raw + n
			if ok := Valid(raw + n); ok {
				luhn = rawn
				return luhn
			}
		}
	}
	luhn = raw
	return luhn
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
