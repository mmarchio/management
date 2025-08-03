{{define "form.prompttemplate"}}
<ul>
    <li><input type="text" name="name" id="name" placeholder="name" value="{{if .Name}}{{.Name}}{{end}}"></li>
    <li><textarea name="template" id="template" placeholder="template">{{if .Template}}{{.Template}}{{end}}</textarea></li>
    <li><textarea name="vars" id="vars" placeholder="vars">{{if .Vars}}{{.Vars}}{{end}}</textarea></li>
    <li>{{template "element.submit" .}}</li>
</ul>
{{end}}