package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Interval2Ints(t *testing.T) {
	ast := assert.New(t)

	actual := Interval2Ints(Interval{Start: 1, End: 2})
	expected := []int{1, 2}
	ast.Equal(expected, actual)
}

func Test_IntervalSlice2Intss(t *testing.T) {
	ast := assert.New(t)

	actual := IntervalSlice2Intss(
		[]Interval{
			Interval{
				Start: 1,
				End:   2,
			},
			Interval{
				Start: 3,
				End:   4,
			},
		},
	)
	expected := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}

	ast.Equal(expected, actual)
}
func Test_Intss2IntervalSlice(t *testing.T) {
	ast := assert.New(t)

	expected := []Interval{
		Interval{
			Start: 1,
			End:   2,
		},
		Interval{
			Start: 3,
			End:   4,
		},
	}
	actual := Intss2IntervalSlice([][]int{
		[]int{1, 2},
		[]int{3, 4},
	},
	)

	ast.Equal(expected, actual)
}
