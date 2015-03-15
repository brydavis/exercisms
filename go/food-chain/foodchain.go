package foodchain

import "strings"

const TestVersion = 1

func Verse(v int) string {
	return ref[v]
}

func Verses(v1, v2 int) string {
	return ref[v1] + "\n\n" + ref[v2]

}

func Song() string {
	return strings.Join(ref[1:], "\n\n")
}
