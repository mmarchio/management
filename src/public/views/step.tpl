{{define "form.step"}}
<ul>
    <li><input type="text" name="name" id="name" placeholder="name" value="{{if .Name}}{{.Name}}{{end}}"></li>
    <li><input type="number" name="order" id="order" placeholder="order" value="{{if .Order}}{{.Order}}{{end}}"></li>
    <li>
        {{$disposition_id := ""}}
        {{if .DispositionID}}{{$disposition_id = .DispositionID}}{{end}}
        <select name="disposition_id" id="disposition_id">
            <option value="">Select Disposition</option>
        {{range .Dispositions}}
            {{if eq $disposition_id .ID}}
            <option value="{{.ID}}" selected="selected">{{.Name}}</option>
            {{else}}
            <option value="{{.ID}}">{{.Name}}</option>
            {{end}}
        {{end}}
        </select>
    </li>
    <li>
        {{$workflow_id := ""}}
        {{if .WorkflowID}}{{$workflow_id = .WorkflowID}}{{end}}
        <select name="workflow_id" id="workflow_id">
            <option value="">Select Workflow</option>
        {{range .Workflows}}
            <option value="{{.ID}}"{{if eq .ID $workflow_id}} selected="selected"{{end}}>{{.Name}}</option>
        {{end}}
        </select>
    </li>
    <li>
        {{$system_prompt_id := ""}}
        {{if .SystemPrompt}}{{$system_prompt_id = .SystemPrompt}}{{end}}
        <select name="system_prompt" id="system_prompt">
            <option value="">Select System Prompt</option>
            {{range .SystemPrompts}}
            <option value="{{.ID}}"{{if eq .ID $system_prompt_id}} selected="selected"{{end}}>{{.Name}}</option>
            {{end}}
        </select>
    </li>
    <li>
        {{$prompt_template_id := ""}}
        {{if .PromptTemplate}}{{$prompt_template_id = .PromptTemplate}}{{end}}
        <select name="prompt_template" id="prompt_template">
            <option value="">Select Prompt Template</option>
            {{range .PromptTemplates}}
            <option value="{{.ID}}"{{if eq .ID $prompt_template_id}} selected="selected"{{end}}>{{.Name}}</option>
            {{end}}
        </select>
    </li>
    <li>
        {{$node_id := ""}}
        {{if .Node}}{{$node_id = .Node}}{{end}}
        <select name="node" id="node">
            <option value="">Select Node</option>
            {{range .Nodes}}
            <option value="{{.ID}}"{{if eq .ID $node_id}} selected="selected"{{end}}>{{.Name}}</option>
            {{end}}
        </select>
    </li>
    <li>{{template "element.toggle" .Enabled}}</li>
    <li>{{template "element.toggle" .Bypass}}</li>
    <li>{{template "element.submit" .}}</li>
</ul>
{{end}}

<!DOCTYPE html>
<html>
    <!--Template head /-->
    {{template "head" .}}
    <body>
        <!-- Template menu /-->
        {{template "menu" .}}
        <!-- Template submenu /-->
        {{template "submenu" .Menu}}
        {{if .DisplayType}}
            {{if eq .DisplayType "none"}}
            <!-- displaytype none /-->
            {{end}}
            {{if eq .DisplayType "new"}}
            <!-- displaytype new /-->
        <form action="/step/save" method="POST" name="new">
        <!-- template form.step /-->
        {{template "form.step" .}}
        </form>
            {{end}}
            {{if eq .DisplayType "edit"}}
        <form action="/step/save{{if .ID}}/{{.ID}}{{end}}" method="POST" name="edit">
        <!-- template form.step /-->
        {{template "form.step" .}}
        </form>
            {{end}}
            {{if eq .DisplayType "list"}}
            <!-- displaytype list /-->
        <table>
            <tr>
                <th>name</th>
                <th>order</th>
                <th>disposition</th>
                <th>node</th>
                <th>system prompt</th>
                <th>prompt template</th>
                <th>enabled</th>
                <th>bypass</th>
            </tr>
            {{range .List}}
            {{$disposition_model_name := ""}}
            {{if .DispositionModel}}{{$disposition_model_name = .DispositionModel.Name}}{{end}}
            {{$workflow_model_name := ""}}
            {{if .WorkflowModel}}{{$workflow_model_name = .WorkflowModel.Name}}{{end}}
            {{$system_prompt_model_name := ""}}
            {{if .NodeModel.SystemPrompt}}{{$system_prompt_model_name = .NodeModel.SystemPrompt}}{{end}}
            {{$prompt_template_model_name := ""}}
            {{if .NodeModel.PromptTemplate}}{{$prompt_template_model_name = .NodeModel.PromptTemplate}}{{end}}
            {{$node_model_name := ""}}
            <tr>
                <td><a href="/step/edit{{if .Step.Model.ID}}/{{.Step.Model.ID}}{{end}}">{{if .Step.Name}}{{.Step.Name}}{{end}}</a></td>
                <td>{{.Order}}</td>
                <td>{{$disposition_model_name}}</td>
                <td><a href="/node/edit/{{.NodeModel.ID}}">{{.NodeModel.Name}}</a></td>
                <td>{{$system_prompt_model_name}}</td>
                <td>{{$prompt_template_model_name}}</td>
                <td>{{if .Enabled.Value}}true{{else}}false{{end}}</td>
                <td>{{if .Bypass.Value}}true{{else}}false{{end}}</td>
                <td><a href="/step/delete{{if .ID}}/{{.ID}}{{end}}">delete</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
    </body>
</html>