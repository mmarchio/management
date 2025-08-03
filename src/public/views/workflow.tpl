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
        <form action="/workflow/save{{if .ID}}/{{.ID}}{{end}}" method="POST">
        {{template "form.workflow" .}}
        </form>
            {{end}}
            {{if eq .DisplayType "list"}}
        <table>
            <tr>
                <th>name</th>
                <th>run</th>
                <th>edit</th>
                <th>remove</th>
            </tr>
            {{range .List}}
            <tr>
                <td>{{.Name}}</td>
                <td><a href="/workflow/run/{{.ID}}">run</a></td>
                <td><a href="/workflow/edit/{{.ID}}">edit</a></td>
                <td><a href="/workflow/delete/{{.ID}}">delete</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
    </body>
</html>