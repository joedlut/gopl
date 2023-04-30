package main

import (
	"ch4/github"
	"html/template"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var issuelist = template.Must(template.New("issuelist").Parse(`
    <h1>{{.TotalCount}}issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
<th>URL</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
<td>{{.State}}</td>
<td>{{.User.Login}}</td>
<td>{{.Title}}</td>
<td>{{.HTMLURL}}</td>
</tr>
{{end}}
</table>
`))
	if err := issuelist.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
