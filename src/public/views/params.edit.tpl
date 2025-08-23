<!DOCTYPE html>
<html>
    {{template "head" .}}
    <body>
        {{template "menu" .}}
        {{template "submenu" .}}
        <form action="/params/save/{{.ID}}" method="POST">
            <input type="hidden" name="node_id" id="node_id" value="{{.ID}}">
            <ul>
                <li><input type="text" name="name" id="name" placeholder="name" value="{{if .Params.Name}}{{.Params.Name}}{{end}}"></li>
                {{if eq .Type "ollama_node"}}
                <li><input type="text" name="ollama_model" id="ollama_model" placeholder="model" value="{{if .Params.OllamaModel}}{{.Params.OllamaModel}}{{end}}"></li>
                <li>
                    <select name="system_prompt">
                    {{range .SystemPrompts}}
                        <option value="{{.ID}}">{{.Name}}:{{.Domain}}</option>
                    {{end}}
                    </select>
                </li>
                <li><input type="text" name="prompt" id="prompt" placeholder="prompt" value="{{if .Prompt.Prompt}}{{.Prompt.Prompt}}{{end}}"></li>
                <li>
                    <select name="prompt_template" id="prompt_template">
                        <option value="">Select Template</option>
                    {{range .PromptTemplates}}
                        {{.}}
                        <option value="{{.ID}}">{{.Name}}</option>
                    {{end}}
                    </select>
                </li>
                {{end}}
        {{if eq .Type "comfy_node"}}
                <li><textarea name="api_base" id="api_base" placeholder="api base">{{if .Params.APIBase}}{{.Params.APIBase}}{{end}}</textarea></li>
                <li><textarea name="api_template" id="api_template" placeholder="api template">{{if .Params.APITemplate}}{{.Params.APITemplate}}{{end}}</textarea></li>
        {{end}}
        {{if eq .Type "ssh_node"}}
                <li><input type="text" name="command" id="command" placeholder="command" value="{{if .Params.Command}}{{.Params.Command}}{{end}}"></li>
                <li><input type="text" name="user" id="user" placeholder="user" value="n8n"></li>
                <li><input type="text" name="host" id="host" placeholder="host" value="172.17.0.1"></li>
        {{end}}
                <li>{{template "element.submit" .}}</li>
            </ul>
        </form>
    </body>
</html>