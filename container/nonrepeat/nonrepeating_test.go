package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"acccbacb", 3},
		{"pwwkew", 3},

		//Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbbbbbbb", 1},
		{"abcadcabcd", 4},

		// Chinese support
		{"这里是地球这里是中华人民共和国这里是北京这里是myway", 12},
		{"一二三四五六五四三二一", 6},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; "+
				"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "这里是地球这里是中华人民共和国这里是北京这里是myway"
	for i := 0; i < 13; i++ {
		s = s + s
	}

	ans := 12

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d",
				actual, s, ans)
		}

	}

}
