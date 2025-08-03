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