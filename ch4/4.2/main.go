package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

func main() {
	sha := flag.String("sh", "sha256", "请输入哈希算法")
	str := flag.String("s", "x", "请输入要加密的字符串")
	flag.Parse()

	fmt.Printf("%s, %s\n", *sha, *str)
	printHash(*sha, *str)
}

func printHash(sha, str string)  {
	switch strings.ToLower(sha) {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
	default:
		fmt.Println("指定加密算法有误")
	}
}
