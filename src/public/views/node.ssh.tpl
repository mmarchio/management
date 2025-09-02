{{define "form.sshnode"}}
               <ul>
                    <li><input type="text" name="name" id="name" value="{{if .Name}}{{.Name}}{{end}}" placeholder="name"></li>
                    <li><input type="text" name="command" id="command" placeholder="prompt" value="{{if .Command}}{{.Command}}{{end}}"></li>
                    <li><input type="text" name="user" id="user" placeholder="user" value="{{if .User}}{{.User}}{{end}}"></li>
                    <li><input type="text" name="host" id="host" placeholder="host" value="{{if .Host}}{{.Host}}{{else}}172.17.0.1{{end}}"></li>
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
            <form action="/node/ssh/save{{if .WorkflowID}}/{{.WorkflowID}}{{end}}" method="post" name="sshnode" id="sshnode">
            {{template "form.sshnode" .}}
            </form>
            {{end}}
            {{if eq .DisplayType "edit"}}
            <form action="/node/ssh/save/{{if .ID}}{{.ID}}{{end}}" method="post" name="sshnode" id="sshnode">
            {{template "form.sshnode" .}}
            </form>
            {{end}}
            {{if eq .DisplayType "list"}}
            {{end}}
        {{end}}
    </body>
</html>