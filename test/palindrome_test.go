package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {

	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},
	}
	for _, test := range tests {
		if test.want != IsPalindrome(test.input) {
			t.Errorf("IsPalindrome(%q) = %v", test.input, !test.want)
		}
	}

}

func TestIsPalindrome2(t *testing.T) {
	seed := time.Now().Unix()
	t.Logf("%d\n", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		t.Log(p)

		if !IsPalindrome(p) {
			t.Error(p)
		}
	}
	if seed == 0 {
		fmt.Println("dfsa")
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}
