package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(plain string) (result string) {
	// Remove all spaces and punctuation
	r := regexp.MustCompile("[^0-9A-Za-z_]")
	// Lowercase all characters
	plain = r.ReplaceAllString(strings.ToLower(plain), "")
	// Measure length of string
	length_of_plain := len(plain)
	// Find the square root of the lenght of string
	sqrt := math.Sqrt(float64(length_of_plain))
	// Assign length of rows and columns
	x := int(sqrt) - 1 // number of horizontal columns
	y := x             // number of vertical rows

	// If not a perfect square...
	if rem := math.Mod(sqrt, 1); rem != 0 {
		z := int(sqrt + 1)
		zz := z * z
		z = (zz - length_of_plain)

		for i := 0; i < z; i++ {
			plain += " "
		}

		y += 1
		x += 1
	}

	// Create crypto square
	for i := 0; i <= y; i++ { // vertical
		for j := 0; j <= x; j++ { // horizontal
			e := j + (y * j) + i
			a := string(plain[e])
			result += a
		}
		if i != y {
			result += " "
		}
	}
	r = regexp.MustCompile("[ ]+")
	result = strings.Trim(r.ReplaceAllString(result, " "), " ")
	return result
}
