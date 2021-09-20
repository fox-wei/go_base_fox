package eratosthenes

import "math"

func Eratosthenes(n int) []int {
	isPrime := make([]bool, n+1) //*筛选质数，默认值false

	for i := 2; i < int(math.Pow(float64(n), 0.5))+1; i++ {
		if !isPrime[i] {
			for j := i * i; j <= (n + 1); j += i {
				isPrime[j] = true
			}
		}
	}

	res := []int{} //?保存结果

	for i := 2; i < n+1; i++ {
		if !isPrime[i] {
			res = append(res, i)
		}
	}

	return res
}
