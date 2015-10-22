package main

import (
	"fmt"
	"math"
)

func main() {
	// f1()
	// f2()
	f3()
}

func f2() {
	fmt.Println(math.Floor(12.1))
}

func f1() {
	// 基数
	var baseNum, keyE, keyD, msg float64
	baseNum = 3 * 11
	// 公钥
	keyE = 3
	// 密钥
	keyD = 7
	// 明文
	msg = 2
	// 加密后的数据
	var encodeMsg = rsa(baseNum, keyE, msg)
	// 解密后的数据
	var decodeMsg = rsa(baseNum, keyD, float64(encodeMsg))
	fmt.Println(encodeMsg)
	fmt.Println(decodeMsg)
}

func rsa(baseNum float64, key float64, message float64) int {
	if baseNum < 1 || key < 1 {
		return 0
	}
	return int(math.Floor(math.Pow(message, key))) % int(baseNum)
}

func f3() {
	// 1、假设p = 3、q = 11（p，q都是素数即可。），则N = pq = 33；
	// 2、r = (p-1)(q-1) = (3-1)(11-1) = 20；
	// 3、根据模反元素的计算公式，我们可以得出，e·d ≡ 1 (mod 20),即e·d = 20n+1 (n为正整数)；
	//    我们假设n=1，则e·d = 21。e、d为正整数，并且e与r互质，则e = 3，d = 7。（两个数交换一下也可以。）
	// 4、到这里，公钥和密钥已经确定。公钥为(N, e) = (33, 3)，密钥为(N, d) = (33, 7)。

	var p, q, N, r int
	p = 19
	q = 23

	N = p * q
	r = (p - 1) * (q - 1)
	fmt.Println(N)
	fmt.Println(r)

}
