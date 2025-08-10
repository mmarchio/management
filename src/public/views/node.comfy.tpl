{{define "form.comfynode"}}
               <ul>
                    <li><input type="text" name="name" id="name" value="" placeholder="name"></li>
                    <li><textarea name="prompt" id="prompt" placeholder="prompt"></textarea></li>
                    <li><textarea name="api_base" id="api_base" placeholder="api base"></textarea></li>
                    <li><textarea name="api_template" id="api_template" placeholder="api template"></textarea></li>
                    <li><textarea name="template_values" id="template_values" placeholder="template values"></textarea></li>
                    <li>{{template "element.toggle" .Enabled}}</li>
                    <li>{{template "element.toggle" .Bypass}}</li>
                    <li><textarea name="output" id="output" value="" placeholder="output"></textarea></li>
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
            <form action="/node/comfy/save" method="post">
            {{template "form.comfynode" .}}
            </form>
            {{end}}
            {{if eq .DisplayType "edit"}}
            <form action="/node/comfy/save/{{if .ID}}{{.ID}}{{end}}" method="post">
            {{template "form.comfynode" .}}
            </form>
            {{end}}
            {{if eq .DisplayType "list"}}
            {{end}}
        {{end}}
    </body>
</html>
