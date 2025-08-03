{{define "form.systemprompts"}}
<ul>
    <li><input type="text" name="name" id="name" value="{{if .Name}}{{.Name}}{{end}}" placeholder="name"></li>
    <li><input type="text" name="domain" id="domain" value="{{if .Domain}}{{.Domain}}{{end}}" placeholder="domain"></li>
    <li><textarea name="prompt" id="prompt" placeholder="prompt">{{if .Prompt}}{{.Prompt}}{{end}}</textarea>
    <li><input type="submit" value="submit"></li>
</ul>
{{end}}