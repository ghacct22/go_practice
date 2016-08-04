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

func main() {
	divs := make([]int, 10000)
	for i := 1; i < 10000; i++ {
		divs[i] = d(i)
	}

	amicables := make(map[int]bool)
	amicableSum := 0
	for i := 1; i < 10000; i++ {
		j := divs[i]
		if j < 10000 && i != j && divs[j] == i {
			amicables[i] = true
			amicables[j] = true
		}
	}

	for key := range amicables {
		amicableSum += key
	}

	fmt.Println(amicableSum)

}
