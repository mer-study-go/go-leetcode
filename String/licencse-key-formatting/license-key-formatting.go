package main

import (
	"fmt"
	"strings"
)

func main() {

	// Example 1
	input1S := "5F3Z-2e-9-w"
	input1K := 4
	fmt.Println(licenseKeyFormatting(input1S, input1K))

	// Example 2
	input2S := "2-5g-3-J"
	input2K := 2
	fmt.Println(licenseKeyFormatting(input2S, input2K))
}

func licenseKeyFormatting(S string, K int) string {

	S = strings.ToUpper(strings.Replace(S, "-", "", -1))
	for i := len(S) - 1 - K; i >= 0; i = i - K {
		S = S[:i+1] + "-" + S[i+1:]
	}
	return S
}
