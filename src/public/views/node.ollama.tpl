{{define "form.ollamanode"}}
                <input type="hidden" name="workflow_id" id="workflow_id" value="{{if .WorkflowID}}{{.WorkflowID}}{{end}}">
                <ul>
                    <li><input type="text" name="name" id="name" value="{{if .OllamaNode.Name}}{{.OllamaNode.Name}}{{end}}" placeholder="name"></li>
                    <li><input type="text" name="ollama_model" id="ollama_model" value="{{if .OllamaModel}}{{.OllamaModel}}{{end}}" placeholder="model"></li>
                    <li>
                        <select name="system_prompt" id="system_prompt">
                            <option value="">Select System Prompt</option>
                            {{$sp := ""}}
                            {{if .SystemPrompt}}{{$sp = .SystemPrompt}}{{end}}
                            {{range .SystemPrompts}}
                            <option value="{{.ID}}"{{if eq .ID $sp}} selected="selected"{{end}}>{{.Name}}:{{.Domain}}</option>
                            {{end}}
                        </select>
                    </li>
                    <li><textarea name="prompt" id="prompt" placeholder="prompt">{{if .Prompt}}{{.Prompt}}{{end}}</textarea></li>
                    <li>
                        {{$pt := ""}}
                        {{if .PromptTemplate}}{{$pt = .PromptTemplate}}{{end}}
                        <select name="prompt_template" id="prompt_template">
                            <option value="">Select Prompt Template</option>
                            {{range .PromptTemplates}}
                            <option value="{{.ID}}"{{if eq $pt .ID}} selected="selected"{{end}}>{{.Name}}</option>
                            {{end}}
                        </select>
                    </li>
                    <li>{{template "element.toggle" .Enabled}}</li>
                    <li>{{template "element.toggle" .Bypass}}</li>
                    <!--<li><textarea name="output" id="output" value="" placeholder="output"></textarea></li> /-->
                    <li>{{template "element.submit" .}}</li>
                </ul>
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
            <form action="/node/ollama/save" method="post">
            {{template "form.ollamanode" .}}
            </form>
            {{end}}
            {{if eq .DisplayType "edit"}}

            {{template "element.debug" .MSI}}

            <form action="/node/ollama/save{{if .ID}}/{{.ID}}{{end}}" method="post" id="ollamanode" name="ollamanode">
            {{template "form.ollamanode" .}}
            </form>
            {{end}}
            {{if eq .DisplayType "list"}}
            {{end}}
        {{end}}
    </body>
</html>
