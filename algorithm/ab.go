package main

import (
	"fmt"
)

func main() {
	for i := 97; i < 97+26; i++ {
		fmt.Printf("%d, %s \n", i, string(rune(i)))
	}
	fmt.Println(rune('吴'))
}
