package main

import (
	"crypto/sha256"
	"fmt"
)



func main() {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))


	fmt.Println(bitCountFunc1(a, b))

	fmt.Println(bitCountFunc2(a, b))
}

func bitCountFunc1(a, b [32]byte) (bitCount uint)  {
	var pc = func() (pc [256]byte) {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i & 1)
		}
		return
	}()

	for i := 0; i < len(a); i++ {
		temp := a[i] ^ b[i]
		bitCount += uint(pc[temp])
	}
	return
}

func bitCountFunc2(a, b [32]byte) (bitCount uint) {
	for i := 0; i < len(a); i++ {
		for j := uint(0); j < 8; j++ {
			if (a[i] & (1 << j)) != (b[i] & (1 << j)) {
				bitCount++
			}
		}
	}
	return
}
