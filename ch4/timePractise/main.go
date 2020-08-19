package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	no1st := time.Date(now.Year(), now.Month(), 1, 0, 0,0,0, time.Local)
	fmt.Println(no1st.AddDate(0, 0, -1))

	beforeOneDay := now.AddDate(0, 0, -1)
	fmt.Println(beforeOneDay)

	afterOneDay := now.AddDate(0, 0, 1)
	fmt.Println(afterOneDay)

	beforeOneMonth := now.AddDate(0, -1, 0)
	afterOneMonth := now.AddDate(0, 1, 0)

	fmt.Println(beforeOneMonth, afterOneMonth)
}
