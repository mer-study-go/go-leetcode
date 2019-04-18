package findpivotindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 7, 3, 6, 5, 6},
		3,
	},
	{
		[]int{1, 2, 3},
		-1,
	},
	{
		[]int{-1, -1, -1, -1, -1, -1},
		-1,
	},
	{
		[]int{-1, -1, -1, 0, 1, 1},
		0,
	},
}

func Test_pivotIndex(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, pivotIndex(tc.nums))
	}
}
