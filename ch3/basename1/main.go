package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename1("a/b/c.go"))
	fmt.Println(basename1("c.d.go"))
	fmt.Println(basename1("abc"))

	fmt.Println(basename2("a/b/c.go"))
	fmt.Println(basename2("c.d.go"))
	fmt.Println(basename2("abc"))
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == '.' {
			s = s[:j]
			break
		}
	}

	return s
}

func basename2(s string) string {
	lastIndexSlash := strings.LastIndex(s, "/")

	if lastIndexSlash != -1 {
		s = s[lastIndexSlash+1:]
	}

	lastIndexDot := strings.LastIndex(s, ".")

	if lastIndexDot != -1 {
		s = s[:lastIndexDot]
	}

	return s
}
