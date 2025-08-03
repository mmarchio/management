
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
            {{if and .Prompt.ID .Prompt.Name}}<span>created: {{.Prompt.ID}}</span>{{end}}
        <form action="/prompts/save{{if .Prompt.ID}}/{{.Prompt.ID}}{{end}}" method="POST">
            <ul>
                <input type="hidden" name="id" id="id" value="{{if .Prompt.Model.ID}}{{.Prompt.Model.ID}}{{end}}">
                <input type="hidden" name="createdAt" id="createdAt" value="{{if .Prompt.Model.CreatedAt}}{{.Prompt.Model.CreatedAt}}{{end}}">
                <input type="hidden" name="updatedAt" id="updatedAt" value="{{if .Prompt.Model.UpdatedAt}}{{.Prompt.Model.UpdatedAt}}{{end}}">
                <input type="hidden" name="contentType" id="contentType" value="{{if .Prompt.Model.ContentType}}{{.Prompt.Model.ContentType}}{{end}}">
                <li><input type="text" name="name" id="name" placeholder="name" value="{{if .Prompt.Name}}{{.Prompt.Name}}{{end}}">
                <li><input type="text" name="prompt" id="prompt" placeholder="prompt" value="{{if .Prompt.Prompt}}{{.Prompt.Prompt}}{{end}}">
                <li><input type="text" name="domain" id="domain" placeholder="domain" value="{{if .Prompt.Domain}}{{.Prompt.Domain}}{{end}}">
                <li><input type="text" name="category" id="category" placeholder="category" value="{{if .Prompt.Category}}{{.Prompt.Category}}{{end}}">
                <li>{{if .Settings}}{{template "form.settings" .}}{{end}}</li>
                <li><input type="submit" value="submit">
            </ul>
        </form>
            {{end}}
            {{if eq .DisplayType "list"}}
        <table>
            <tr>
                <th>Name</th>
                <th>Prompt</th>
                <th>Domain</th>
                <th>Category</th>
                <th>Settings</th>
                <th>Delete</th>
            </tr>
            {{range $prompt := .List}}
            <tr>
                <td><a href="/prompts/edit/{{$prompt.ID}}">{{$prompt.Name}}</a></td>
                <td>{{$prompt.Prompt}}</td>
                <td>{{$prompt.Domain}}</td>
                <td>{{$prompt.Category}}</td>
                <td><a href="/prompts/settings/{{$prompt.ID}}">edit</a></td>
                <td><a href="/prompts/delete/{{$prompt.ID}}">delete</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
    </body>
</html>