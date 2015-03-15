package secret

import "strconv"

var binmap = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

func Handshake(n int) []string {
	var store []string

	if n > 0 && n < 32 {
		a := int64(n)
		b := ReverseString(strconv.FormatInt(a, 2))

		for i, v := range b {
			if i < 4 {
				if "1" == string(v) {
					store = append(store, binmap[i])
				}
			} else if 4 == i && "1" == string(v) {
				return ReverseSlice(store)
			}
		}

	}
	return store
}

func ReverseString(input string) string {
	var revString string

	for _, r := range input {
		revString = string(r) + revString
	}

	return revString
}

func ReverseSlice(s []string) []string {
	n := len(s)
	revSlice := make([]string, n)

	for i, r := range s {
		revSlice[n-1-i] = r
	}

	return revSlice
}
