package raindrops

import "fmt"

func Convert(num int) string {
	output := ""
	if num%3 == 0 {
		output += "Pling"
	}
	if num%5 == 0 {
		output += "Plang"
	}
	if num%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = fmt.Sprintf("%d", num)
	}
	return output
}
