package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	fmt.Println("hello")
	f1()
}

func f1() {

	priv, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	fmt.Println(priv.N.BitLen())
}

func f2() {

}
