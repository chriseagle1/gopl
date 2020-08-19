package github

import "time"

const IssuesURL = "https://api.github.com/search/code"

type IssuesSearchResult struct {
	TotalCount int
	Iterms []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}
