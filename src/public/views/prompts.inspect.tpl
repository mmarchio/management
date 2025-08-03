<!DOCTYPE html>
<html>
    {{template "head" .}}
    <body>
        {{template "menu" .}}
        {{template "submenu" .Menu}}
            {{if .ID}}<span>created: {{.ID}}</span>{{end}}
            <span style="float:left; clear:both;">{{if .content.Content}}{{.content.Content}}{{end}}</span>
    </body>
</html>