package decrease

//?快乘法，根据俄罗斯农民问题实现
func QuickMulti(a, b int) int {
	ans := 0
	for ; a > 0; a >>= 1 {
		if (a & 1) == 1 {
			ans += b
		}
		b <<= 1
	}
	return ans
}

func RussionMulti(a, b int) int {
	ans := 0
	if a == 0 || b == 0 {
		return 0
	}

	for a > 1 {
		if (a % 2) == 0 {
			a /= 2
		} else {
			a /= 2
			ans += b
		}
		b *= 2
	}

	return ans + b
}
