package sieve

import "math"

func Sieve(n int) []int {
	var primeList = []int{2}
	var isPrime int = 1
	var num int = 3
	var sqrtNum int = 0
	for primeList[len(primeList)-1] <= n && num <= n {
		sqrtNum = int(math.Sqrt(float64(num))) // Sqrt(num)
		for i := 0; i < len(primeList); i++ {
			if num%primeList[i] == 0 {
				isPrime = 0
			}
			if primeList[i] > sqrtNum {
				i = len(primeList)
			}
		}
		if isPrime == 1 {
			primeList = append(primeList, num)
		} else {
			isPrime = 1
		}
		num = num + 2
	}
	return primeList
}
