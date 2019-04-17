package arrangecoins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase suite slice
var tcs = []struct {
	n   int
	ans int
}{
	{
		1804289383,
		60070,
	},
	{
		10,
		4,
	},
	{
		5,
		2,
	},
	{
		8,
		3,
	},

	// More testcases go here
}

func Test_arrangeCoins(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, arrangeCoins(tc.n), "Input: %v", tc)
	}
}
