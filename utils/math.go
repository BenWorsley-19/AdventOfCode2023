package utils

import "log"

func LeastCommonMultiple(numbers []int64) int64 {
	if len(numbers) < 2 {
		log.Fatal("lcm requires at least 2 numbers")
	}

	var result int64
	for i := 0; i < len(numbers)-1; i++ {
		a := result
		if i == 0 {
			a = numbers[i]
		}
		b := numbers[i+1]
		gcd := greatestCommonDivisor(a, b)
		result = a * b / gcd
	}
	return result
}

// Euclidean Algorithm
func greatestCommonDivisor(a int64, b int64) int64 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	r := a % b
	return greatestCommonDivisor(b, r)
}
