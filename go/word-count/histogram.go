package wc

import (
	"regexp"
	"strings"
)

type Histogram map[string]int

func WordCount(text string) Histogram {
	hist := Histogram{}
	rgx := regexp.MustCompile(`[a-zA-Z|0-9]+`)
	chopped := rgx.FindAllString(strings.ToLower(text), -1)
	for _, v := range chopped {
		if hist[v] > 0 {
			hist[v] += 1
		} else {
			hist[v] = 1
		}
	}
	return hist
}
