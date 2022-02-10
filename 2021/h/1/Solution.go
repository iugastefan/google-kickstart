package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 1; i <= T; i++ {
		var S, F string
		n, err := fmt.Scan(&S, &F)
		if n != 2 || err != nil {
			fmt.Println("Error at test case", i)
			panic(err)
		}
		fmt.Printf("Case #%d: %d\n", i, solve(S, F))
	}

}

func solve(S, F string) int {
	charMap := make(map[rune]rune)
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	for _, i := range F {
		for _, c := range alphabet {
			dist := distance(i, c)
			old, ok := charMap[c]
			if ok {
				if dist < distance(old, c) {
					charMap[c] = i
				}
			} else {
				charMap[c] = i
			}

		}
	}
	sum := 0
	for _, i := range S {
		sum += distance(i, charMap[i])
	}
	return sum
}

func distance(a, b rune) int {
	intA, intB := float64(a), float64(b)
	return int(math.Min(math.Abs(intA-intB), 26-math.Abs(intB-intA)))
}
