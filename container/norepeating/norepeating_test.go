package main

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcadefc", 6},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"aaaaaa", 1},
		{"b", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"一二三二一", 3},
		{"这里是慕课网", 6},
		{"黑化肥会发会发灰灰化肥挥发会发黑", 6},
	}

	for _, tt := range tests {
		actual := lengthOfNoneRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d.", actual, tt.s, tt.ans)
		}
	}
}

// cd .
// go test(go test .)
// go test -coverprofile=c.out
// go tool cover -html=c.out

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥会发会发灰灰化肥挥发会发黑"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s)=%d", len(s))
	b.ResetTimer()

	ans := 6

	for i := 0; i < b.N; i++ {
		actual := lengthOfNoneRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d.", actual, s, ans)
		}
	}
}

// go test -bench .
// go test -bench . -cpuprofile cpu.out
// go tool pprof cpu.out
// help
// web
