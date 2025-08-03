<!DOCTYPE html>
<html>
    {{template "head" .}}
    <body>
        {{template "menu" .}}
        {{template "submenu" .Menu}}
        {{if .DisplayType}}
            {{if eq .DisplayType "none"}}
            {{end}}
            {{if eq .DisplayType "new"}}
        <form action="/node/save{{if .ID}}/{{.ID}}{{end}}" method="POST">
        {{template "form.node" .}}
        </form>
            {{end}}
            {{if eq .DisplayType "list"}}
        <table>
            <tr>
                <th>name</th>
                <th>type</th>
                <th>enabled</th>
                <th>bypass</th>
                <th>output</th>
                <th>run</th>
                <th>edit</th>
                <th>delete</th>
            </tr>
            {{range .List}}
            <tr>
                <td>{{.Name}}</td>
                <td>{{.Type}}</td>
                <td>{{if .Enabled}}true{{else}}false{{end}}</td>
                <td>{{if .Bypass}}true{{else}}false{{end}}</td>
                <td>{{.Output}}</td>
                <td><a href="/node/run/{{.ID}}">run</a></td>
                <td><a href="/node/edit/{{.ID}}">edit</a></td>
                <td><a href="/node/delete/{{.ID}}">delete</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
    </body>
</html>