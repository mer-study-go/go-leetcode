package main

import "fmt"

func main() {
	A := []int{12, 28, 46, 32, 50}
	B := []int{50, 12, 32, 46, 28}
	fmt.Println(anagramMappings(A, B))
}

func anagramMappings(A []int, B []int) []int {
	dict := make(map[int]int)
	for i := 0; i < len(B); i++ {
		dict[B[i]] = i
	}
	var result []int
	for j := 0; j < len(A); j++ {
		result = append(result, dict[A[j]])
	}

	return result
}
