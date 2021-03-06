package eratosthenes

import (
	"math"
	"reflect"
)

/*
*Sieve of Eratosthenes 埃拉托斯特尼筛法
* 求n以内的全部的素数
*核心原理：素数的倍数一定是合数
 */

func Eratosthenes(n int) []int {
	isPrime := make([]bool, n+1) //*筛选素数，默认值false

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

func AllSame(parm ...interface{}) bool {
	arr := reflect.ValueOf(parm[0])
	v := arr.Index(0).Interface()

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() != v {
			return false
		}
	}
	return true
}
