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
        {{range .OllamaNodes}}
        <tr>
                <td>{{.Name}}</td>
                <td>{{.Type}}</td>
                <td><a href="/ollama/node/edit/{{.ID}}">edit</a></td>
                <td><a href="/ollama/node/delete/{{.ID}}">delete</a></td>
        </tr>
        {{end}}
        {{range .SSHNodes}}
        <tr>
                <td>{{.Name}}</td>
                <td>{{.Type}}</td>
                <td><a href="/ssh/node/edit/{{.ID}}">edit</a></td>
                <td><a href="/ssh/node/delete/{{.ID}}">delete</a></td>
        </tr>
        {{end}}
        {{range .ComfyNodes}}
        <tr>
                <td>{{.Name}}</td>
                <td>{{.Type}}</td>
                <td><a href="/comfy/node/edit/{{.ID}}">edit</a></td>
                <td><a href="/comfy/node/delete/{{.ID}}">delete</a></td>
        </tr>
        {{end}}
</table>
{{end}}