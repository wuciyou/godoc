package main

import (
	"fmt"
)

var baseArr []int

func main() {
	baseArr = []int{10, 2, 3, 53, 21, 31, 54, 64, 9, 6, 2, 65, 100}
	// fmt.Println(baseArr)

	fmt.Println(kuaisu(baseArr, 0, len(baseArr)-1))
}

// 快速排序
func kuaisu(baseArr []int, left int, right int) []int {
	if baseArr == nil || len(baseArr) <= 1 {
		return baseArr
	}
	var baseVal = baseArr[0]

	var leftArr []int
	var rightArr []int

	for left < right {
		// var leftVal int
		var rightVal int

		// 先从右往左找
		for ; right >= left; right-- {
			rightVal = baseArr[right]
			if rightVal < baseVal {
				leftArr = append(leftArr, rightVal)
				break
			}
		}
		fmt.Println(leftArr)
	}

	return append(leftArr, rightArr...)
	// fmt.Println(baseArr)
}
