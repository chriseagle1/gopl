package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	//reverse(a[:])
	//fmt.Println(a)

	reverse2(&a)

	fmt.Println(a)
}

//func reverse(s []int) {
//	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
//		s[i], s[j] = s[j], s[i]
//	}
//}

func reverse2(p *[6]int)  {
	for i, j := 0, len(p) - 1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}