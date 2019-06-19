package main

import (
	"math"
)

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}

	flag := math.Sqrt(float64(n))
	for i := 2; i <= int(flag) + 1; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func Compute(data_chunk int) (result Result) {
	for i := 0; i < data_chunk; i++ {
		var primes []int
		for number := 1; number <= 10; number++ {
			if IsPrime(number) {
				primes = append(primes, number)
			}
		}
		result = append(result, primes)
	}
	return result
}
