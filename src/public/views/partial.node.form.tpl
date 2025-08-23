{{define "form.node"}}
<ul>
        <input type="hidden" name="workflow_id" id="workflow_id" value="{{if .WorkflowID}}{{.WorkflowID}}{{end}}">
        <li><input type="text" name="name" id="name" placeholder="name" value="{{if .Name}}{{.Name}}{{end}}"></li>
        <li>
                <select name="type" id="type">
                        <option value="">select type</option>
                        <option value="ollama_node">Ollama Node</option>
                        <option value="comfy_node">Comfy Node</option>
                        <option value="ssh_node">SSH Node</option>
                </select>
        </li>
        <li><span>Enabled</span><label class="switch"><input type="checkbox" name="enabled" id="enabled" {{if eq .Enabled true}}checked{{end}}><span class="slider round"></span></label></li>
        <li><span>Bypass</span><label class="switch"><input type="checkbox" name="bypass" id="bypass" {{if eq .Bypass true}}checked{{end}}><span class="slider round"></span></label></li>
        <li><textarea name="output" id="output" placeholder="output">{{if .Output}}{{.Output}}{{end}}</textarea></li>
        <li>{{template "element.submit" .}}</li>
</ul>
{{end}}