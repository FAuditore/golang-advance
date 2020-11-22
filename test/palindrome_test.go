package test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}

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
//go test -bench=BenchmarkIsPalindrome -cpuprofile=cpu.out -benchmem -covermode=count -coverprofile=cover.out
//go tool cover -html=cover.out
//go tool pprof cpu.out
//web