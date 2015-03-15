package hamming

import "errors"

const TestVersion = 1

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) == len(b) {
		for k, v := range a {
			if v != rune(b[k]) {
				count += 1
			}
		}
		return count, nil
	}
	return 0, errors.New("Strand lengths vary...")
}
