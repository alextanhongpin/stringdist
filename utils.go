package stringdist

func min(head int, rest ...int) int {
	for _, i := range rest {
		if i < head {
			head = i
		}
	}
	return head
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
