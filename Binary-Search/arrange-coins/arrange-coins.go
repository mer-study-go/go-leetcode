package arrangecoins

func arrangeCoins(n int) int {
	return helper(n, 1, n)
}

func helper(n, start, end int) int {

	for start+1 < end {
		mid := start + (end-start)/2
		if sum(mid) <= n {
			start = mid
		} else {
			end = mid
		}
	}

	if sum(end) == n {
		return end
	}

	return start
}

func sum(n int) int {
	return n * (n + 1) / 2
}
