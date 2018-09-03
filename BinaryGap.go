func Solution(N int) int {
	// write your code in Go 1.4
	maxGap, currentGap, isFirstOne := 0, 0, true
	for N != 0 {
		if N&1 == 1 {
			if currentGap > maxGap {
				maxGap = currentGap
			}
			currentGap = 0
			isFirstOne = false
		} else {
			if !isFirstOne {
				currentGap++
			}
		}
		N >>= 1
	}
	return maxGap
}