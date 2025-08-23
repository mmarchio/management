{{define "form.comfy"}}
<ul>
    <li><input type="text" name="name" id="name" value="{{if .Name}}{{.Name}}{{end}}" placeholder="name"></li>
    <li><input type="text" name="endpoint" id="endpoint" value="{{if .Endpoint}}{{.Endpoint}}{{end}}" placeholder="endpoint"></li>
    <li><textarea name="base" id="base" placeholder="base">{{if .Base}}{{.Base}}{{end}}</textarea></li>
    <li><textarea name="template" id="template" placeholder="template">{{if .Template}}{{.Template}}{{end}}</textarea></li>
    <li><input type="submit" value="submit"></li>
</ul>
{{end}}