package main

import "fmt"

var pc [256]byte = func() (pc [256]byte){
	for i := range pc  {
		pc[i] = pc[i/2] + byte(i & 1)
	}
	return
}()



func main() {
	//var x uint64 = 127   //     01111111
	//start := time.Now()
	//PopCount(x)
	//dura1 := time.Since(start).Seconds()
	//fmt.Printf("%.10f\n", dura1)
	//
	//start2 := time.Now()
	//PopCountWithLoop(x)
	//dura2 := time.Since(start2).Seconds()
	//fmt.Printf("%.10f\n", dura2)

	x := "hello!"

	for i := 0; i < len(x) ; i++ {
		x := x[i]

		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}

//func PopCount(x uint64) int  {
//	return int(pc[byte(x >> (0 * 8))] +
//		pc[byte(x >> (1 * 8))] +
//		pc[byte(x >> (2 * 8))] +
//		pc[byte(x >> (3 * 8))] +
//		pc[byte(x >> (4 * 8))] +
//		pc[byte(x >> (5 * 8))] +
//		pc[byte(x >> (6 * 8))] +
//		pc[byte(x >> (7 * 8))])
//}
//
//func PopCountWithLoop(x uint64) int  {
//	var tmp byte
//	var i uint8
//
//	for i = 0; i < 8; i++ {
//		tmp += pc[byte(x >> (i * 8))]
//	}
//
//	return int(tmp)
//}

