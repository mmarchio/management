{{define "form.comfy"}}
<ul>
    <li><input type="text" name="name" id="name" value="{{if .Name}}{{.Name}}{{end}}" placeholder="name"></li>
    <li><input type="text" name="endpoint" id="endpoint" value="{{if .Endpoint}}{{.Endpoint}}{{end}}" placeholder="endpoint"></li>
    <li><textarea name="base" id="base" placeholder="base">{{if .Base}}{{.Base}}{{end}}</textarea></li>
    <li><textarea name="template" id="template" placeholder="template">{{if .Template}}{{.Template}}{{end}}</textarea></li>
    <li><input type="submit" value="submit"></li>
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
            {{if eq .DisplayType "list"}}
            {{end}}
            {{if eq .DisplayType "new"}}
        <form action="/comfy/save{{if .ID}}/{{.ID}}{{end}}" method="POST">
        {{template "form.comfy" .}}
        </form>
            {{end}}
        {{end}}
    </body>
</html>