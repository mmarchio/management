{{define "form.template"}}
<h2>template</h2>
<details>
<summary>template</summary>
<ul>
    <li><input type="text" name="template_name" id="template_name" placeholder="name" value="{{if .Name}}{{.Name}}{{end}}"></li>
    {{range .Dispositions}}
    <li><a href="/disposition/edit/{{.ID}}">edit {{.Name}}</a></li>
    {{end}}
    <li>
        {{range .Dispositions}}
        <li><span style="float:left">{{.Name}}</span><input style="float:right" type="checkbox" name="{{.ID}}" id="{{.ID}}" checked></li>
        {{end}}
        {{range .AvailableDispositions}}
        <li><span style="float:left">{{.Name}}</span><input style="float:right" type="checkbox" name="{{.ID}}" id="{{.ID}}"></li>
        {{end}}
    </li>
    <li><a href="/dispositions/new">new disposition</a></li>
    <li><input type="number" name="template_current_disposition" id="template_current_disposition" placeholder="current disposition" value="{{if .CurrentDisposition}}{{.CurrentDisposition}}{{end}}"></li>
</ul>
</details>
{{end}}