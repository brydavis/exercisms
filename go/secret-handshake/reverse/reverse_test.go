package reverse

import (
	"testing"
)

const (
	s       = "The quick brown 狐 jumped over the lazy 犬"
	reverse = "犬 yzal eht revo depmuj 狐 nworb kciuq ehT"
)

func TestReverse(t *testing.T) {
	if RevTwo(s) != reverse {
		t.Error(s)
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RevTwo(s)
	}
}
