package main

import (
	"fmt"
	"math"
)

func main() {
	// f3(10)
	f4(12)
}

func f4(n int) {
	var sun, sun1, sun2 int
	sun1, sun2 = 1, 1
	sun = sun1 + sun2
	// 1 1 2 3 5 8 13 21 34 55
	for i := 2; i < n; i++ {
		sun += sun1 + sun2
		sun1 = sun2
		sun2 = sun
	}
	fmt.Println(sun)
}

func f3(n int) {
	var laugh, laugh1 int
	// (n/2)-1
	for i := 1; i <= n; i *= 2 {
		laugh1++
		for j := 1; j <= i; j++ {
			laugh++
		}
	}
	fmt.Printf("laugh:%d, laugh1:%d \n", laugh, laugh1)
}

func f1(num int) int {
	var counts int
	for i := 1; i <= num; i++ {
		counts += i
	}
	return counts
}

func f2(num int) int {
	return int(float64(num+1) * math.Floor(float64(num)/2))
}
