{{define "form.systemprompts"}}
<ul>
    <li><input type="text" name="name" id="name" value="{{if .Name}}{{.Name}}{{end}}" placeholder="name"></li>
    <li><input type="text" name="self_model" id="self_model" value="{{if .SelfModel}}{{.SelfModel}}{{end}}" placeholder="self model"></li>
    <li><input type="text" name="target_model" id="target_model" value="{{if .TargetModel}}{{.TargetModel}}{{end}}" placeholder="target model"></li>
    <li><input type="text" name="domain" id="domain" value="{{if .Domain}}{{.Domain}}{{end}}" placeholder="domain"></li>
    <li><textarea name="prompt" id="prompt" placeholder="prompt">{{if .Prompt}}{{.Prompt}}{{end}}</textarea>
    <li><input type="submit" value="submit"></li>
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
        <form action="/systemprompts/save{{if .ID}}/{{.ID}}{{end}}" method="POST">
        {{template "form.systemprompts" .}}
        </form>
            {{end}}
            {{if eq .DisplayType "list"}}
        <table>
            <tr>
                <th>id</th>
                <th>name</th>
                <th>domain</th>
                <th>prompt</th>
                <th>delete</th>
            </tr>
            {{range $systemprompt := .List}}
            <tr>
                <td><a href="/systemprompts/{{$systemprompt.ID}}">{{$systemprompt.ID}}</a></td>
                <td>{{$systemprompt.Name}}</td>
                <td>{{$systemprompt.Domain}}</td>
                <td>{{$systemprompt.Prompt}}</td>
                <td><a href="/systemprompts/delete/{{$systemprompt.ID}}">delete</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
    </body>
</html>