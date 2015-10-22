package main

import (
	"fmt"
	"strings"
)

func main() {
	// trimfunc("abcdefg")
	ContainsAny("[cidyou]")
	ContainsAny("cidyou")
}

func ContainsAny(str string) {
	fmt.Println(strings.ContainsAny(str, "[&]"))
}

func trimfunc(str string) {
	fmt.Println(str)
	str = strings.TrimFunc(str, func(r rune) bool {
		if r == 'g' {
			return true
		}
		return false
	})
	fmt.Println(str)
}
