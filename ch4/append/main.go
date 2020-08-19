package main

import "fmt"

func main() {
	var x  = []int{1, 2}

	fmt.Println(appendInt(x, 3))

	m := make([]int, 0, 3)
	fmt.Printf("%x\t%d\t%d\n", m, len(m), cap(m))

	m = appendInt(m, 2)
	fmt.Printf("%x\t%d\t%d\n", m, len(m), cap(m))

	m = appendInt(m, 5)
	fmt.Printf("%x\t%d\t%d\n", m, len(m), cap(m))

	m = appendInt(m, 6)
	fmt.Printf("%x\t%d\t%d\n", m, len(m), cap(m))

	m = appendInt(m, 8)
	fmt.Printf("%x\t%d\t%d\n", m, len(m), cap(m))

	m = appendInt(m, 3)
	fmt.Printf("%x\t%d\t%d\n", m, len(m), cap(m))
}

func appendInt(x []int, y int) []int  {
	var z []int

	zlen := len(x) + 1

	if zlen <= cap(x){
		z = x[:zlen]
	} else {
		zcap := zlen

		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y

	return z
}
