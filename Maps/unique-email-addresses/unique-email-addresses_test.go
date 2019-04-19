package uniqueemailaddresses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input []string
	ans   int
}{
	{
		[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
		2,
	},
	{
		[]string{"testemail@leetcode.com", "testemail1@leetcode.com", "testemail+david@lee.tcode.com"},
		3,
	},
}

func Test_numUniqueEmails(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numUniqueEmails(tc.input))
	}
}
