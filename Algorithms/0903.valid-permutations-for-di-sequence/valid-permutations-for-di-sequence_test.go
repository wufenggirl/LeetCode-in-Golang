package problem0903

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans int
}{

	{
		"DID",
		5,
	},

	{
		"DIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDDDIDD",
		795732224,
	},

	// 可以有多个 testcase
}

func Test_numPermsDISequence(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numPermsDISequence(tc.S), "输入:%v", tc)
	}
}

func Benchmark_numPermsDISequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numPermsDISequence(tc.S)
		}
	}
}
