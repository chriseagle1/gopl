package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1234567"))
	fmt.Println(comma("123"))
	fmt.Println(comma("12"))

	fmt.Println(comma2("1234567"))
	fmt.Println(comma2("123"))
	fmt.Println(comma2("12"))
	fmt.Println(comma2("1.2"))
	fmt.Println(comma2("12.579"))
	fmt.Println(comma2("125743396.334"))
	fmt.Println(comma2("12572443396.334998"))

	fmt.Println(compareStr("1a2b3c", "1a2b3c"))
	fmt.Println(compareStr("1a2b3c", "b3c1a2"))
	fmt.Println(compareStr("1a2b3c", "1a2d3c"))
}

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}


func comma2(s string) string {
	var buf bytes.Buffer

	dot := strings.Index(s, ".")

	var intStr string

	if dot > 0 {
		intStr = s[:dot]
	} else {
		intStr = s
	}

	step := len(intStr) % 3

	if step == 0 {
		step = 3
	}

	buf.WriteString(s[:step])

	for step < len(intStr) {
		buf.WriteByte(',')
		buf.WriteString(s[step:step+3])
		step += 3
	}
	if dot > 0 {
		buf.WriteString(s[dot:])
	}

	return buf.String()
}

func compareStr(s1, s2 string) bool  {
	if strings.Compare(s1, s2) == 0 {
		return false
	}

	for _, ch := range s1 {
		if !strings.ContainsRune(s2, ch) {
			return false
		}
	}

	return true
}