package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	go func() {
		for {

		}
	}()
	go func() {
		println("here")
	}()

	for {
	}
}
