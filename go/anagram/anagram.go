package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) (result []string) {
	subject = strings.ToLower(subject)
	reversed := SortString(subject)
	for _, v := range candidates {
		v = strings.ToLower(v)
		candidate := SortString(v)
		if reversed == candidate && v != subject {
			result = append(result, v)
		}
	}
	return
}

func SortString(text string) string {
	var sorted []string
	for _, v := range text {
		sorted = append(sorted, string(v))
	}
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
