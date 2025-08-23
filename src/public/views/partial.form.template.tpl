{{define "form.template"}}

<h2>template</h2>
<details>
<summary>template</summary>
<ul>
    <li><input type="text" name="template_name" id="template_name" placeholder="name" value="{{if .Name}}{{.Name}}{{end}}"></li>
    <li>
        <ul>
        {{range .AvailableDispositions}}
        <li><span style="float:left">{{if .Name}}{{.Name}}{{end}}</span><input style="float:right" type="checkbox" name="{{if .ID}}{{.ID}}{{end}}" id="{{if .ID}}{{.ID}}{{end}}"></li>
        {{end}}
        {{range .DispositionsArrayModel}}
        <li><span style="float:left">{{if .Name}}{{.Name}}{{end}}</span><input style="float:right" type="checkbox" name="{{if .ID}}{{.ID}}{{end}}" id="{{if .ID}}{{.ID}}{{end}}" checked></li>
        {{end}}
        </ul>
    </li>
    <li><a href="/dispositions/new">new disposition</a></li>
    <li><input type="number" name="template_current_disposition" id="template_current_disposition" placeholder="current disposition" value="{{if .CurrentDisposition}}{{.CurrentDisposition}}{{end}}"></li>
</ul>
</details>
{{end}}