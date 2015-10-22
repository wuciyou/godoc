package main

import (
	"fmt"
	"hash/crc32"
	"time"
)

func main() {
	// crc32s()
	// fmt.Println((1 << 31))

	// fmt.Println(math.MaxInt32)
	for i := 0; i <= 10000000; i++ {
		times()
	}

}

func times() {
	start := time.Now().UnixNano()
	end := time.Now().UnixNano()
	run := end - start
	if run > 50 {
		fmt.Println(run)
	}

}

func crc32s() {

	hs := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	hs.Write([]byte("wuciyoudfs"))
	s := hs.Sum32()
	fmt.Println(s)
}
