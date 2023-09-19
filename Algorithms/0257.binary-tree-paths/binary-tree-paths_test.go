package problem0257

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     []string
}{

	{
		[]int{},
		[]int{},
		nil,
	},

	{
		[]int{1, 2, 5, 3},
		[]int{2, 5, 1, 3},
		[]string{"1->2->5", "1->3"},
	},

	// 可以有多个 testcase
}

func Test_binaryTreePaths(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.ans, binaryTreePaths(root), "输入:%v", tc)
	}
}

func Benchmark_binaryTreePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			binaryTreePaths(root)
		}
	}
}
