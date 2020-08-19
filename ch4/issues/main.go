package main

import (
	"awesomeProject/gopl/ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()

	beforeOneMonth := now.AddDate(0, -1, 0)
	beforeOneYear := now.AddDate(-1, 0, 0)

	var postMonth []*github.Issues
	var postYear []*github.Issues
	var postMoreYear []*github.Issues

	for _, item := range result.Items{
		if item.CreateAt.Before(beforeOneYear) {
			postMoreYear = append(postMoreYear, item)
		} else if item.CreateAt.Before(beforeOneMonth) {
			postYear = append(postYear, item)
		} else {
			postMonth = append(postMonth, item)
		}
	}

	fmt.Printf("一个月内的有%d\n", len(postMonth))
	for _, item := range postMonth{
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User, item.Title)
	}

	fmt.Printf("一年内的有%d\n", len(postYear))
	for _, item := range postYear{
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User, item.Title)
	}

	fmt.Printf("超过一年的有%d\n", len(postMoreYear))
	for _, item := range postMoreYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User, item.Title)
	}
}
