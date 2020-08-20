package main

import (
	"awesomeProject/gopl/ch4/github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .items}}------------------------------
Number:{{.Number}}
User: {{.User.Login}}
Title:{{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

var report = template.Must(template.New("issueList").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours()/24)
}