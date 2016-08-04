package main

import "fmt"

func divisors(n int) []int {
	var divs []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
		}
	}
	return divs
}

func d(n int) int {
	sum := 0
	for _, div := range divisors(n) {
		sum += div
	}
	return sum
}

func abundantCheck(abs []int) map[int]bool {
	result := make(map[int]bool)
	for _, x := range abs {
		for _, y := range abs {
			result[x+y] = true
		}
	}
	return result
}

func main() {
	var abundants []int
	for i := 12; i < 28123; i++ {
		if d(i) > i {
			abundants = append(abundants, i)
		}
	}

	abundantSum := 0
	abundantSums := abundantCheck(abundants)
	fmt.Println(abundantSums)
	for i := 1; i < 28123; i++ {
		if !abundantSums[i] {
			abundantSum += i
		}
	}

	fmt.Println(abundantSum)
}
