package main

import (
	"awesomeProject/gopl/ch2/tempconv"
	"fmt"
)

func main() {
	boilingF := tempconv.CToF(tempconv.FreezingC)

	fmt.Println(tempconv.BoilingC - tempconv.FreezingC)

	fmt.Println(boilingF - tempconv.CToF(tempconv.FreezingC))

	fmt.Println(boilingF.String())
	fmt.Printf("%v\n", boilingF)

	fmt.Printf("%s\n", boilingF)

	fmt.Println(boilingF)

	fmt.Printf("%g\n", boilingF)

	fmt.Println(float64(boilingF))
}
