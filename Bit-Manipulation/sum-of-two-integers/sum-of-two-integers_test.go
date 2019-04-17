package sumoftwointegers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase suite slice
var tcs = []struct {
	a   int
	b   int
	ans int
}{
	{
		1,
		2,
		3,
	},
	{
		-2,
		3,
		1,
	},
}

func Test_getSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, getSum(tc.a, tc.b))
	}
}
