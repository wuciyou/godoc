package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 110; i++ {
		fmt.Printf("\033[0;%dm(%d)", i, i)
		if i%10 == 0 {
			fmt.Println("")
		}
	}

}
