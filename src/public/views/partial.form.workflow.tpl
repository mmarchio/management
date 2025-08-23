{{define "form.workflow"}}
<ul>
        <li><input type="text" name="name" id="name" placeholder="name" value="{{if .Name}}{{.Name}}{{end}}">
        <li><a href="/node/new{{if .ID}}/{{.ID}}{{end}}">New Node</a></li>
        <li>{{template "element.submit" .}}</li>
</ul>
<table style="float:right; clear: none; margin-right: 1000px;">
        <tr>
                <th>name</th>
                <th>type</th>
                <th>params</th>
                <th>edit</th>
                <th>delete</th>
        </tr>
        {{range .Nodes}}
        <tr>
                <td>{{.Name}}</td>
                <td>{{.Type}}</td>
                <td><a href="/params/edit/{{.ID}}">params</a></td>
                <td><a href="/node/edit/{{.ID}}">edit</a></td>
                <td><a href="/node/delete/{{.ID}}">delete</a></td>
        </tr>
        {{end}}
</table>
{{end}}