package issuehtml

import "html/template"

var issuelist = template.Must(template.New("issuelist").Parse(`
 <h1>{{.TotalCount}}issues</h1>
<table>
<tr style='text-align:left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}</td>
<td>{{.State}}</td>
<td>{{.User.Login}}</td>
<td>{{.Title}}</td>
</tr>
{{end}}
</table>
`))
