{{define "form.prompttemplate"}}
<ul>
    <li><input type="text" name="name" id="name" placeholder="name" value="{{if .Name}}{{.Name}}{{end}}"></li>
    <li><textarea name="template" id="template" placeholder="template">{{if .Template}}{{.Template}}{{end}}</textarea></li>
    <li><textarea name="vars" id="vars" placeholder="vars">{{if .Vars}}{{.Vars}}{{end}}</textarea></li>
    <li>{{template "element.submit" .}}</li>
</ul>
{{end}}
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
        <form action="/prompttemplates/save{{if .ID}}/{{.ID}}{{end}}" method="POST" style="float:left; clear:left;">
        {{template "form.prompttemplate" .}}
        </form>
        <pre style="float:right; margin-right: 1000px;">
        {{.Context}}
        </pre>
            {{end}}
            {{if eq .DisplayType "edit"}}
        <form action="/prompttemplates/save{{if .ID}}/{{.ID}}{{end}}" method="POST" style="float:left; clear:left;">
        {{template "form.prompttemplate" .}}
        </form>
        <pre style="float:right; margin-right: 1000px;">
        {{.Context}}
        </pre>
            {{end}}
            {{if eq .DisplayType "list"}}
        <table>
            <tr>
                <th>name</th>
                <th>template</th>
                <th>vars</th>
                <th>delete</th>
            </tr>
            {{range .List}}
            <tr>
                <td><a href="/prompttemplates/edit/{{.ID}}">{{.Name}}</a></td>
                <td>{{.Template}}</td>
                <td>{{.Vars}}</td>
                <td><a href="/prompttemplates/delete/{{.ID}}">delete</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
    </body>
</html>