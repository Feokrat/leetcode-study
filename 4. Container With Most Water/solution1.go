package ContainerWithMostWater

// Time Limit Exceeded

func maxArea(height []int) int {
	n := len(height)
	maxSquare := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			m := min(height[i], height[j])
			square := (j - i) * m
			maxSquare = max(square, maxSquare)
		}
	}
	return maxSquare
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}