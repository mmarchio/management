{{define "form.workflow"}}
<ul>
        <li><input type="text" name="name" id="name" placeholder="name" value="{{if .Name}}{{.Name}}{{end}}">
        <li><a href="/node/comfy/new{{if .ID}}/{{.ID}}{{end}}">New ComfyUI Node</a></li>
        <li><a href="/node/ollama/new{{if .ID}}/{{.ID}}{{end}}">New Ollama Node</a></li>
        <li><a href="/node/ssh/new{{if .ID}}/{{.ID}}{{end}}">New SSH Node</a></li>
        <li>{{template "element.submit" .}}</li>
</ul>
<table style="float:right; clear: none; margin-right: 1000px;">
        <tr>
                <th>name</th>
                <th>type</th>
                <th>edit</th>
                <th>delete</th>
        </tr>
        {{range .OllamaNodesArrayModel}}
        <tr>
                <td>{{.Name}}</td>
                <td>Ollama Node</td>
                <td><a href="/node/ollama/edit/{{.ID}}">edit</a></td>
                <td><a href="/node/ollama/delete/{{.ID}}">delete</a></td>
        </tr>
        {{end}}
        {{range .SSHNodesArrayModel}}
        <tr>
                <td>{{.Name}}</td>
                <td>SSH Node</td>
                <td><a href="/node/ssh/edit/{{.ID}}">edit</a></td>
                <td><a href="/node/ssh/delete/{{.ID}}">delete</a></td>
        </tr>
        {{end}}
        {{range .ComfyNodesArrayModel}}
        <tr>
                <td>{{.Name}}</td>
                <td>ComfyUI Node</td>
                <td><a href="/node/comfy/edit/{{.ID}}">edit</a></td>
                <td><a href="/node/comfy/delete/{{.ID}}">delete</a></td>
        </tr>
        {{end}}
</table>
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
        <form action="/workflow/save{{if .ID}}/{{.ID}}{{end}}" method="POST" name="workflow" id="workflow">
        {{template "form.workflow" .}}
        </form>
            {{end}}
            {{if eq .DisplayType "edit"}}
        <form action="/workflow/save{{if .ID}}/{{.ID}}{{end}}" method="POST" name="workflow" id="workflow">
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