{{define "form.comfynode"}}
                {{if .ComfyNode.WorkflowID}}<input type="hidden" name="workflow_id" value="{{.ComfyNode.WorkflowID}}">{{end}}
               <ul>
                    <li><input type="text" name="name" id="name" value="{{if .ComfyNode.Name}}{{.ComfyNode.Name}}{{end}}" placeholder="name"></li>
                    <li><textarea name="prompt" id="prompt" placeholder="prompt">{{if .ComfyNode.Prompt}}{{.ComfyNode.Prompt}}{{end}}</textarea></li>
                    <li><textarea name="api_base" id="api_base" placeholder="api base">{{if .ComfyNode.APIBase}}{{.ComfyNode.APIBase}}{{end}}</textarea></li>
                    <li><textarea name="api_template" id="api_template" placeholder="api template">{{if .ComfyNode.APITemplate}}{{.ComfyNode.APITemplate}}{{end}}</textarea></li>
                    <li><textarea name="template_values" id="template_values" placeholder="template values">{{if .ComfyNode.TemplateValues}}{{.ComfyNode.TemplateValues}}{{end}}</textarea></li>
                    <li>
                        <select name="service" id="service">
                            <option value="">select service</option>
                            {{$serviceid := 0}}
                            {{if .ServicePort}}{{$serviceid = .ServicePort}}{{end}}
                            {{range .Services}}
                            <option value="{{.Port}}" {{if eq $serviceid .Port}} selected="selected"{{end}}>{{.Title}}</option>
                            {{end}}
                        </select>
                    </li>
                    <li>{{template "element.toggle" .Enabled}}</li>
                    <li>{{template "element.toggle" .Bypass}}</li>
                    <li><textarea name="output" id="output" value="" placeholder="output"></textarea></li>
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
            <form action="/node/comfy/save{{if .WorkflowID}}/{{.WorkflowID}}{{end}}" method="post" name="comfynode" id="comfynode">
            {{template "form.comfynode" .}}
            </form>
            {{end}}
            {{if eq .DisplayType "edit"}}
            <form action="/node/comfy/save{{if .ID}}/{{.ID}}{{end}}" method="post" name="comfynode" id="comfynode">
            {{template "form.comfynode" .}}
            </form>
            {{end}}
            {{if eq .DisplayType "list"}}
            <table>
                <tr>
                    <th>name</th>
                    <th>enabled</th>
                    <th>bypass</th>
                    <th>delete</th>
                </tr>
                {{range .List}}
                <tr>
                    <td><a href="/node/comfy/edit/{{.ID}}">{{.Name}}</a></td>
                    <td>{{if .Enabled.Value}}true{{else}}false{{end}}</td>
                    <td>{{if .Bypass.Value}}true{{else}}false{{end}}</td>
                    <td><a href="/node/comfy/delete/{{.ID}}">delete</a></td>
                </tr>
                {{end}}
            </table>
            {{end}}
        {{end}}
    </body>
</html>
