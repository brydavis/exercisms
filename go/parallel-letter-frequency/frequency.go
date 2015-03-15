package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(slice []string) FreqMap {
	m, channel, payload := FreqMap{}, make(chan rune), len(slice)
	for _, text := range slice {
		go func(text string) {
			for _, rn := range text {
				channel <- rn
			}
			payload -= 1
		}(text)
	}
	for payload > 0 {
		m[<-channel]++
	}
	return m
}
