package slice

import "fmt"

// All returns a list of all substrings of s with length n.
// func All(n int, s string) []string {
// 	for i, _ := range s {
// 		string(s[i : i+n])
// 	}

// }

func All(n int, s string) []string {
	var store []string
	l := len(s)
	for i, _ := range s {
		if (i + n) <= l {
			// fmt.Println(string(s[i : i+n]))
			store = append(store, string(s[i:i+n]))
		}
	}
	return store
}

// Frist returns the first substring of s with length n.
func Frist(n int, s string) string {
	a := All(n, s)
	return a[0]
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
